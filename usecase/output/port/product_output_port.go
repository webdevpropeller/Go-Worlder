package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProductOutputPort ...
type ProductOutputPort interface {
	Index([]model.Product) []outputdata.Product
	Show(*model.Product) *outputdata.Product
	Edit(*model.Product) *outputdata.Product
}
