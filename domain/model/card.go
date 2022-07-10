package model

import (
	"go_worlder_system/validator"
	"time"

	log "github.com/sirupsen/logrus"
)

type Card struct {
	ID           string
	OwnerID      string
	Company      string
	Name         string
	Number       string `validate:"required,len=16"`
	Expiry       *time.Time
	SecurityCode string `validate:"required,min=3,max=4"`
}

func NewCard(
	owner *User,
	company string,
	name string,
	number string,
	expiry *time.Time,
	securityCode string,
) (*Card, error) {
	id := generateID(idPrefix.CardID)
	card := &Card{
		ID:           id,
		OwnerID:      owner.ID,
		Company:      company,
		Name:         name,
		Number:       number,
		Expiry:       expiry,
		SecurityCode: securityCode,
	}
	err := validator.Struct(card)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return card, nil
}
