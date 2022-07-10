package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ChatPresenter ...
type ChatPresenter struct {
}

// NewChatPresenter ...
func NewChatPresenter() *ChatPresenter {
	return &ChatPresenter{}
}

func (presenter *ChatPresenter) IndexDestination(userList []model.Profile) []outputdata.PublicUser {
	oPublicUserList := []outputdata.PublicUser{}
	for _, user := range userList {
		oPublicUser := outputdata.PublicUser{
			ID:   user.UserID,
			Name: user.Company,
		}
		oPublicUserList = append(oPublicUserList, oPublicUser)
	}
	return oPublicUserList
}

// Show ...
func (presenter *ChatPresenter) Show(msgList []model.ChatMessage) []outputdata.ChatMessage {
	oMsgList := []outputdata.ChatMessage{}
	for _, msg := range msgList {
		oMsg := presenter.convert(&msg)
		oMsgList = append(oMsgList, *oMsg)
	}
	return oMsgList
}

// Send ...
func (presenter *ChatPresenter) Send(msg *model.ChatMessage) *outputdata.ChatMessage {
	return presenter.convert(msg)
}

func (presenter *ChatPresenter) convert(msg *model.ChatMessage) *outputdata.ChatMessage {
	return &outputdata.ChatMessage{
		ID:       msg.ID,
		Sender:   msg.Sender.ID,
		Receiver: msg.Receiver.ID,
		Content:  msg.Content,
	}
}
