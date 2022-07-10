package model

import (
	"go_worlder_system/errs"
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

const (
	ReceivingType = iota + 1
	ShippingType
	DisposalType
)

// Inventory is an entity
type Inventory struct {
	ID              string   `validate:"required"`
	User            *User    `validate:"required"`
	Product         *Product `validate:"required"`
	InventoryTypeID int      `validate:"min=1,max=3"`
	Quantity        int      `validate:"min=0"`
}

// NewInventory ...
func NewInventory(user *User, product *Product, inventoryTypeID int, quantity int) (*Inventory, error) {
	id := generateID(idPrefix.InventoryID)
	inventory := &Inventory{
		ID:              id,
		User:            user,
		Product:         product,
		InventoryTypeID: inventoryTypeID,
		Quantity:        quantity,
	}
	err := validator.Struct(inventory)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	// Update product inventory
	switch inventoryTypeID {
	case ReceivingType:
		product.Inventory.Receiving += inventory.Quantity
		product.Inventory.Stock += inventory.Quantity
	case ShippingType:
		product.Inventory.Shipping += inventory.Quantity
		product.Inventory.Stock -= inventory.Quantity
	case DisposalType:
		product.Inventory.Disposal += inventory.Quantity
		product.Inventory.Stock -= inventory.Quantity
	default:
		return nil, errs.Invalidated.New("Inventory type is invalid")
	}
	if product.Inventory.Stock < 0 {
		return nil, errs.Conflict.New("The user can't update inventory")
	}
	return inventory, nil
}
