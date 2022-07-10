package repository

import "go_worlder_system/domain/model"

// UserTokenRepository ...
type UserTokenRepository interface {
	FindByToken(string) (*model.UserToken, error)
	Save(*model.UserToken) error
	Delete(*model.UserToken) error
}
