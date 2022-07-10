package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// Cash ...
type Cash float64

// AccountType ...
type AccountType struct {
	ID   int
	Name string
}

// AccountName ...
type AccountName struct {
	ID   string
	Type *AccountType
	Name string
}

// Account is an entity
type Account struct {
	ID      string       `validate:"required"`
	User    *User        `validate:"required"`
	Name    *AccountName `validate:"required"`
	Balance *AccountBalance
}

// AccountBalance ...
type AccountBalance float64

// NewAccount ...
func NewAccount(user *User, accountName *AccountName) (*Account, error) {
	id := generateID(idPrefix.AccountID)
	var balance AccountBalance
	account := &Account{
		ID:      id,
		User:    user,
		Name:    accountName,
		Balance: &balance,
	}
	err := validator.Struct(account)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return account, nil
}

// IsOwner ...
func (account *Account) IsOwner(userID string) bool {
	return account.User.ID == userID
}
