package repository

import "go_worlder_system/domain/model"

// AccountNameRepository ...
type AccountNameRepository interface {
	FindList() ([]model.AccountName, error)
	FindByID(string) (*model.AccountName, error)
}
