package repository

import "go_worlder_system/domain/model"

// BrandRepository is an interface for DB operation
type BrandRepository interface {
	FindListByUserID(userID string) ([]model.Brand, error)
	FindByID(id string) (*model.Brand, error)
	FindListByLikeUserID(id string) ([]model.Brand, error)
	Save(*model.Brand) error
	Update(*model.Brand) error
	Delete(*model.Brand) error
}
