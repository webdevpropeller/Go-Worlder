package controller

import (
	"encoding/json"
	"go_worlder_system/consts"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/gateway/datasource"
	"go_worlder_system/interface/gateway/localcache"
	"go_worlder_system/interface/gateway/remotecache"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	outputdata "go_worlder_system/usecase/output/data"
	"net/http"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

type ChatMessageRequest struct {
	DestinationID string
	Content       string
}

// ChatController ...
type ChatController struct {
	inputport inputport.ChatInputPort
	clients   map[string]*websocket.Conn
	broadcast chan *outputdata.ChatMessage
}

// NewChatController ...
func NewChatController(sqlHandler database.SQLHandler, noSQLHandler database.NoSQLHandler, cache localcache.Cache) *ChatController {
	broadcast := make(chan *outputdata.ChatMessage)
	clients := make(map[string]*websocket.Conn)
	return &ChatController{
		inputport: interactor.NewChatInteractor(
			presenter.NewChatPresenter(),
			database.NewUserDatabase(sqlHandler),
			datasource.NewChatMessageDataSource(
				database.NewChatMessageDatabase(sqlHandler),
				localcache.NewChatMessageLocalCache(cache),
				remotecache.NewChatMessageRemoteCache(),
			),
		),
		clients:   clients,
		broadcast: broadcast,
	}
}

func (controller *ChatController) BroadCast() {
	for msg := range controller.broadcast {
		receiver, ok := controller.clients[msg.Receiver]
		if ok {
			err := receiver.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}
		sender, ok := controller.clients[msg.Sender]
		if ok {
			err := sender.WriteJSON(msg)
			if err != nil {
				log.Error(err)
			}
		}
	}
}

// WebSocket ...
// @summary Get WebSocket
// @description
// @tags Chat
// @accept json
// @produce json
// @param Connection header string true "Upgrade"
// @param Upgrade header string true "websocket"
// @param message body inputdata.ChatMessageRequest true "Receiving request"
// @success 200
// @failure 400
// @router /chat/ws [post]
func (controller *ChatController) WebSocket(c Context) error {
	tokenString := c.QueryParam(pn.Jwt)
	jwtToken, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			errMsg := "unexpected signing method: %v"
			alg := token.Header["alg"]
			log.Errorf(errMsg, alg)
			return "", errs.Unauthorized.Errorf(errMsg, alg)
		}
		return []byte(os.Getenv(consts.JWT_SECRET)), nil
	})
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), "Login required")
		return err
	}
	claims, ok := jwtToken.Claims.(jwt.MapClaims)
	if !ok {
		errMsg := "not found claims in %s"
		log.Errorf(errMsg, tokenString)
		err := errs.Unauthorized.Errorf(errMsg, tokenString)
		c.String(statusCode(err), "Login required")
		return err
	}
	userID, ok := claims["sub"].(string)
	if !ok {
		errMsg := "not found %s in %s"
		log.Errorf(errMsg, "sub", tokenString)
		err = errs.Unauthorized.Errorf("not found %s in %s", "sub", tokenString)
		c.String(statusCode(err), "Login required")
		return err
	}
	if !c.IsWebSocket() {
		errMsg := "Connection isn't websocket"
		log.Error(errMsg)
		return c.String(statusCode(errs.Unknown.New(errMsg)), errMsg)
	}
	ws, err := c.WebSocket()
	if err != nil {
		log.Error(err)
		return c.String(statusCode(err), err.Error())
	}
	controller.clients[userID] = ws
	for {
		var iMsgRequest ChatMessageRequest
		_, message, err := ws.ReadMessage()
		if err != nil {
			log.Error(err)
			break
		}
		err = json.Unmarshal(message, &iMsgRequest)
		if err != nil {
			log.Error(err)
			break
		}
		iMsg := &inputdata.ChatMessage{
			SenderID:   userID,
			ReceiverID: iMsgRequest.DestinationID,
			Content:    iMsgRequest.Content,
		}
		oMsg, err := controller.inputport.Send(iMsg)
		if err != nil {
			log.Error(err)
			c.String(statusCode(err), errs.Cause(err).Error())
			continue
		}
		controller.broadcast <- oMsg
	}
	delete(controller.clients, userID)
	return ws.Close()
}

// IndexDestination ...
// @summary Get destinations
// @description
// @tags Chat
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @success 200
// @failure 400
// @router /chat/destinations [get]
func (controller *ChatController) IndexDestination(c Context) error {
	userID := c.UserID()
	oPublicUserList, err := controller.inputport.IndexDestination(userID)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oPublicUserList)
}

// Show ...
// @summary Get messages
// @description
// @tags Chat
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "messanger id"
// @success 200 {array} outputdata.ChatMessage "Message list"
// @failure 400
// @router /chat/:id [get]
func (controller *ChatController) Show(c Context) error {
	userID := c.UserID()
	messengerID := c.Param(idParam)
	oMsgList, err := controller.inputport.Show(userID, messengerID)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oMsgList)
}
