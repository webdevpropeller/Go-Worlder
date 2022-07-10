package repository

import (
	"go_worlder_system/domain/model"
)

// PayeeRepository ...
type PayeeRepository interface {
	FindListByUserID(userID string) ([]model.Payee, error)
	FindByID(id string) (*model.Payee, error)
	Save(*model.Payee) error
	Update(*model.Payee) error
	Delete(*model.Payee) error
}
