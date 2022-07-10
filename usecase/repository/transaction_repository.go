package repository

import (
	"go_worlder_system/domain/model"
)

// TransactionRepository is an interface for DB operation
type TransactionRepository interface {
	FindListByUserID(userID string) ([]model.Transaction, error)
	FindListByAccountID(accountID string) ([]model.Transaction, error)
	FindByID(id string) (*model.Transaction, error)
	FindAccountItemList() ([]model.AccountItem, error)
	FindAccountItemByID(id string) (*model.AccountItem, error)
	FindPartnerList() ([]model.TransactionPartner, error)
	FindPartnerByID(partnerID string) (*model.TransactionPartner, error)
	FindAccountByID(accountID string) (*model.TransactionAccount, error)
	Save(*model.Transaction) error
	Update(*model.Transaction) error
	Delete(*model.Transaction) error
}
