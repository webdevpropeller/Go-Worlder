package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// BrandInputPort ...
type BrandInputPort interface {
	Index(userID string) ([]outputdata.Brand, error)
	Show(id string) (*outputdata.Brand, error)
	Create(*inputdata.NewBrand) error
	Edit(id string, userID string) (*outputdata.Brand, error)
	Update(*inputdata.UpdatedBrand) error
	Delete(id string, userID string) error
}
