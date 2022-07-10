package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// ChatMessage ...
type ChatMessage struct {
	ID       string `validate:"required"`
	Sender   *User  `validate:"required"`
	Receiver *User  `validate:"required"`
	Content  string `validate:"required"`
}

// NewChatMessage ...
func NewChatMessage(sender, receiver *User, content string) (*ChatMessage, error) {
	id := generateID(idPrefix.MessageID)
	msg := &ChatMessage{
		ID:       id,
		Sender:   sender,
		Receiver: receiver,
		Content:  content,
	}
	err := validator.Struct(msg)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return msg, nil
}
