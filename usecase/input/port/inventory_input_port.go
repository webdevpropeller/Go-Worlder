package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// InventoryInputPort ...
type InventoryInputPort interface {
	IndexProduct(userID string) ([]outputdata.Inventory, error)
	ShowProduct(productID string, userID string) (*outputdata.Inventory, error)
	EditProduct(productID string, userID string) (*outputdata.Inventory, error)
	UpdateProduct(iManagement *inputdata.ProductManagement, userID string) error
	New(userID string) ([]outputdata.Inventory, error)
	Create(*inputdata.Inventory) error
	UpdateStocktaking(*inputdata.Stocktaking) error
}
