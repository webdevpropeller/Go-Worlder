package repository

import "go_worlder_system/domain/model"

// CategoryRepository ...
type CategoryRepository interface {
	FindByID(id uint) (*model.Category, error)
}
