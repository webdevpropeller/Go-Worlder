package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// SearchOutputPort ...
type SearchOutputPort interface {
	SearchProduct([]model.Product) []outputdata.Product
}
