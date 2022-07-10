package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// PayeeInputPort ...
type PayeeInputPort interface {
	Index(userID string) ([]outputdata.Payee, error)
	Create(*inputdata.NewPayee) error
	Update(*inputdata.UpdatedPayee) error
	Delete(id string, userID string) error
}
