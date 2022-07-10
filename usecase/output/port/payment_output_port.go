package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// PaymentOutputPort ...
type PaymentOutputPort interface {
	Key(*model.APIKey) *outputdata.APIKey
}
