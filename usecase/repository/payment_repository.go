package repository

import "go_worlder_system/domain/model"

// PaymentRepository ...
type PaymentRepository interface {
	Create(*model.Payment) error
}
