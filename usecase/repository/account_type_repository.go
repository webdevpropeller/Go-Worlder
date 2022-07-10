package repository

import "go_worlder_system/domain/model"

// AccountTypeRepository ...
type AccountTypeRepository interface {
	FindByID(uint) (*model.AccountType, error)
}
