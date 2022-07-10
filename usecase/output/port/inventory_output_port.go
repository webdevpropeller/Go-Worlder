package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// InventoryOutputPort ...
type InventoryOutputPort interface {
	IndexProduct([]model.Product) []outputdata.Inventory
	ShowProduct(*model.Product) *outputdata.Inventory
	Editproduct(*model.Product) *outputdata.Inventory
	New([]model.Product) []outputdata.Inventory
}
