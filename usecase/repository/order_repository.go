package repository

import "go_worlder_system/domain/model"

// OrderRepository ...
type OrderRepository interface {
	Create(*model.Order) error
}
