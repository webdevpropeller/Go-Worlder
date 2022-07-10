package repository

import (
	"go_worlder_system/domain/model"
)

// InventoryRepository ...
type InventoryRepository interface {
	FindListByUserIDAndTypeID(userID string, inventoryTypeID int) ([]model.Inventory, error)
	Create(*model.Inventory) error
	Update(*model.Inventory) error
	Delete(*model.Inventory) error
}
