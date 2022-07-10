package model

import (
	"go_worlder_system/errs"
	"go_worlder_system/validator"
	"mime/multipart"

	log "github.com/sirupsen/logrus"
)

// Product ...
type Product struct {
	ID          string    `validate:"required"`
	Brand       *Brand    `validate:"required"`
	Category    *Category `validate:"required"`
	GenreID     uint
	Name        string `validate:"required"`
	Price       uint   `validate:"min=1"`
	Image       *multipart.FileHeader
	Description string
	IsDraft     bool
	Inventory   *ProductInventory
	Management  *ProductManagement
}

// ProductInventory ...
type ProductInventory struct {
	ProductID string
	Receiving int `validate:"min=0"`
	Shipping  int `validate:"min=0"`
	Disposal  int `validate:"min=0"`
	Stock     int `validate:"min=0"`
}

// ProductManagement ...
type ProductManagement struct {
	ProductID string
	Storage   string
	Memo      string
	Barcode   string
}

// NewProduct ...
func NewProduct(
	brand *Brand, category *Category, genreID uint,
	name string, price uint, image *multipart.FileHeader,
	description string, isDraft bool,
) (*Product, error) {
	id := generateID(idPrefix.ProductID)
	inventory := &ProductInventory{
		ProductID: id,
		Receiving: 0,
		Shipping:  0,
		Disposal:  0,
		Stock:     0,
	}
	management := &ProductManagement{
		ProductID: id,
		Storage:   "",
		Memo:      "",
		Barcode:   "",
	}
	product := &Product{
		ID:          id,
		Brand:       brand,
		Category:    category,
		GenreID:     genreID,
		Name:        name,
		Price:       price,
		Image:       image,
		Description: description,
		IsDraft:     isDraft,
		Inventory:   inventory,
		Management:  management,
	}
	err := validator.Struct(product)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return product, nil
}

// IsOwner ...
func (product *Product) IsOwner(userID string) bool {
	return product.Brand.User.ID == userID
}

// Update ...
func (product *Product) Update(
	category *Category, genreID uint, name string,
	price uint, image *multipart.FileHeader,
	description string, isDraft bool,
) error {
	if product.IsDraft == false && isDraft == true {
		errMsg := "Product cannot be returned to draft"
		log.Error(errMsg)
		return errs.Conflict.New(errMsg)
	}
	product.Category = category
	product.GenreID = genreID
	product.Name = name
	product.Price = price
	product.Image = image
	product.Description = description
	err := validator.Struct(product)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdateManagement ...
func (product *Product) UpdateManagement(storage string, barcode string, memo string) error {
	product.Management.Storage = storage
	product.Management.Barcode = barcode
	product.Management.Memo = memo
	err := validator.Struct(product.Management)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Stocktaking ...
func (product *Product) Stocktaking(stock int) error {
	product.Inventory.Stock = stock
	if product.Inventory.Stock < 0 {
		return errs.Conflict.New("The user can't update inventory")
	}
	return nil
}
