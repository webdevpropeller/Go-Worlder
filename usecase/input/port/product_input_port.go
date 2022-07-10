package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProductInputPort ...
type ProductInputPort interface {
	Index(userID string) ([]outputdata.Product, error)
	Show(id string) (*outputdata.Product, error)
	Create(*inputdata.NewProduct) error
	Update(*inputdata.UpdatedProduct) error
	Edit(id string, userID string) (*outputdata.Product, error)
	Delete(id string, userID string) error
}
