package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// AccountInputPort ...
type AccountInputPort interface {
	Index(userID string) ([]outputdata.Account, error)
	Show(id string, userID string) (*outputdata.Account, error)
	New() ([]outputdata.AccountName, error)
	Create(*inputdata.Account) error
	Delete(id string, userID string) error
}
