package database

import (
	"go_worlder_system/domain/model"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ChatMessageDatabase ...
type ChatMessageDatabase struct {
	SQLHandler
}

// NewChatMessageDatabase ...
func NewChatMessageDatabase(sqlHandler SQLHandler) *ChatMessageDatabase {
	return &ChatMessageDatabase{sqlHandler}
}

func (db *ChatMessageDatabase) FindDestinationListByUserID(userID string) ([]model.Profile, error) {
	profileTable := userDB.ProfileTable
	chatMessagesTable := chatDB.ChatMessagesTable
	columnList := generateColumnList(profileTable)
	columnQuery := generateSelectColumnQuery(columnList)
	statement := strings.Join([]string{
		rw.Select, rw.Distinct, columnQuery,
		rw.From, profileTable.NAME(),
		rw.RightJoin, chatMessagesTable.NAME(),
		rw.On, profileTable.UserID, "=", chatMessagesTable.ReceiverID,
		rw.Where, chatMessagesTable.SenderID, "= ?",
		rw.Union,
		rw.Select, rw.Distinct, columnQuery,
		rw.From, profileTable.NAME(),
		rw.RightJoin, chatMessagesTable.NAME(),
		rw.On, profileTable.UserID, "=", chatMessagesTable.SenderID,
		rw.Where, chatMessagesTable.ReceiverID, "= ?",
	}, " ")
	rows, err := db.Query(statement, userID, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	userList := []model.Profile{}
	for rows.Next() {
		user := model.Profile{}
		scanList := generateScanList(&user)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		if user.UserID == userID {
			continue
		}
		userList = append(userList, user)
	}
	return userList, nil
}

func (db *ChatMessageDatabase) FindListByUserIDs(user1ID, user2ID string) ([]model.ChatMessage, error) {
	chatMessagesTable := chatDB.ChatMessagesTable
	columnList := generateColumnList(chatMessagesTable)
	columnQuery := generateSelectColumnQuery(columnList)
	statement := strings.Join([]string{
		rw.Select, columnQuery,
		rw.From, chatMessagesTable.NAME(),
		rw.Where,
		chatMessagesTable.SenderID, rw.In, "(?,?)", rw.And,
		chatMessagesTable.ReceiverID, rw.In, "(?,?)",
		rw.OrderBy, chatMessagesTable.CreatedAt,
	}, " ")
	rows, err := db.Query(statement, user1ID, user2ID, user1ID, user2ID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	msgList := []model.ChatMessage{}
	for rows.Next() {
		msg := model.ChatMessage{
			Sender:   &model.User{},
			Receiver: &model.User{},
		}
		scanList := []interface{}{&msg.ID, &msg.Sender.ID, &msg.Receiver.ID, &msg.Content}
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		msgList = append(msgList, msg)
	}
	return msgList, nil
}

// Save ...
func (db *ChatMessageDatabase) Save(msg *model.ChatMessage) error {
	_, err := db.Transaction(func() (interface{}, error) {
		chatMessagesTable := chatDB.ChatMessagesTable
		statement := NewSQLBuilder().Insert(chatMessagesTable)
		valueList := generateValueList(msg)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
