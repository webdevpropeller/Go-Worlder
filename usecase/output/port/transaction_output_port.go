package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// TransactionOutputPort ...
type TransactionOutputPort interface {
	Index([]model.Transaction) []outputdata.Transaction
	Show(*model.Transaction) *outputdata.Transaction
	Edit(*model.Transaction) *outputdata.Transaction
}
