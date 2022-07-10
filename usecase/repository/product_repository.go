package repository

import (
	"go_worlder_system/domain/model"
)

// ProductRepository ...
type ProductRepository interface {
	RetrieveListByKeyWord(keyword string) ([]model.Product, error)
	FindListByUserID(userID string) ([]model.Product, error)
	FindListByBrandID(brandID string) ([]model.Product, error)
	FindByID(id string) (*model.Product, error)
	Create(*model.Product) error
	Update(*model.Product) error
	UpdateInventory(*model.ProductInventory) error
	UpdateManagement(*model.ProductManagement) error
	Delete(id string) error
	DeleteListByBrandID(brandID string) error
}
