package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// AccountOutputPort ...
type AccountOutputPort interface {
	Index([]model.Account) []outputdata.Account
	Show(*model.Account) *outputdata.Account
	New([]model.AccountName) []outputdata.AccountName
}
