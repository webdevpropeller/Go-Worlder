package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// BrandOutputPort ...
type BrandOutputPort interface {
	Index([]model.Brand) []outputdata.Brand
	Show(*model.Brand) *outputdata.Brand
	Edit(*model.Brand) *outputdata.Brand
}
