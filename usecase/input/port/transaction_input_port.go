package inputport

import (
	"go_worlder_system/domain/model"
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// TransactionInputPort ...
type TransactionInputPort interface {
	New() ([]model.AccountItem, error)
	Index(userID string) ([]outputdata.Transaction, error)
	IndexAccount(accountID string, userID string) ([]outputdata.Transaction, error)
	Create(*inputdata.NewTransaction) error
	Edit(id string, userID string) (*outputdata.Transaction, error)
	Update(*inputdata.UpdatedTransaction) error
	Delete(id string, userID string) error
}
