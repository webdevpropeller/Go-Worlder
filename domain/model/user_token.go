package model

import (
	"go_worlder_system/str"
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// UserToken ...
type UserToken struct {
	User  *User  `validate:"required"`
	Token string `validate:"required"`
}

// NewUserToken ...
func NewUserToken(user *User) (*UserToken, error) {
	token := str.Random(tokenLength)
	userToken := &UserToken{
		User:  user,
		Token: token,
	}
	err := validator.Struct(userToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return userToken, nil
}
