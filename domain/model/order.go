package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// Order is an entity
type Order struct {
	ID       string   `validate:"required"`
	User     *User    `validate:"required"`
	Product  *Product `validate:"required"`
	Quantity int
}

// NewOrder ...
func NewOrder(user *User, product *Product, quantity int) (*Order, error) {
	id := generateID(idPrefix.OrderID)
	order := &Order{
		ID:       id,
		User:     user,
		Product:  product,
		Quantity: quantity,
	}
	err := validator.Struct(order)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return order, nil
}
