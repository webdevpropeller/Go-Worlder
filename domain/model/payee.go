package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// Payee ...
type Payee struct {
	ID   string `validate:"required"`
	User *User  `validate:"required"`
	Name string `validate:"required"`
}

// NewPayee ...
func NewPayee(user *User, name string) (*Payee, error) {
	id := generateID(idPrefix.PayeeID)
	payee := &Payee{
		ID:   id,
		User: user,
		Name: name,
	}
	err := validator.Struct(payee)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return payee, nil
}

// IsOwner ...
func (payee *Payee) IsOwner(userID string) bool {
	return payee.User.ID == userID
}
