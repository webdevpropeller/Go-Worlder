package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// PayeeOutputPort ...
type PayeeOutputPort interface {
	Index([]model.Payee) []outputdata.Payee
	Show(*model.Payee) *outputdata.Payee
	Edit(*model.Payee) *outputdata.Payee
}
