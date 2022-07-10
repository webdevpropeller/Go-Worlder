package repository

import "go_worlder_system/domain/model"

// APIKeyRepository ...
type APIKeyRepository interface {
	FindByUserID(string) (*model.APIKey, error)
	Create(model.User) error
}
