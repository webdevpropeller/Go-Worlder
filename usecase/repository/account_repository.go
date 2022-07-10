package repository

import (
	"go_worlder_system/domain/model"
)

// AccountRepository is an interface for DB operation
type AccountRepository interface {
	FindListByUserID(userID string) ([]model.Account, error)
	FindByID(id string) (*model.Account, error)
	FindCashByUserID(userID string) (*model.Cash, error)
	Save(*model.Account) error
	Delete(*model.Account) error
}
