package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// InventoryPresenter ...
type InventoryPresenter struct {
}

// NewInventoryPresenter ...
func NewInventoryPresenter() outputport.InventoryOutputPort {
	return &InventoryPresenter{}
}

// IndexProduct ...
func (presenter *InventoryPresenter) IndexProduct(productList []model.Product) []outputdata.Inventory {
	oInventoryList := []outputdata.Inventory{}
	for _, product := range productList {
		oInventory := presenter.convert(&product)
		oInventoryList = append(oInventoryList, *oInventory)
	}
	return oInventoryList
}

// ShowProduct ...
func (presenter *InventoryPresenter) ShowProduct(product *model.Product) *outputdata.Inventory {
	return presenter.convert(product)
}

// Editproduct ...
func (presenter *InventoryPresenter) Editproduct(product *model.Product) *outputdata.Inventory {
	return presenter.convert(product)
}

// New ...
func (presenter *InventoryPresenter) New(productList []model.Product) []outputdata.Inventory {
	oInventoryList := []outputdata.Inventory{}
	for _, product := range productList {
		oInventory := presenter.convert(&product)
		oInventoryList = append(oInventoryList, *oInventory)
	}
	return oInventoryList
}

func (presenter *InventoryPresenter) convert(product *model.Product) *outputdata.Inventory {
	oUser := &outputdata.UserSimplified{
		ID:   product.Brand.User.ID,
		Name: product.Brand.User.Profile.Company,
	}
	oBrand := &outputdata.BrandSimplified{
		ID:   product.Brand.ID,
		Name: product.Brand.Name,
	}
	oCategory := &outputdata.Category{
		ID:   product.Category.ID,
		Name: product.Category.Name,
	}
	oProduct := &outputdata.Product{
		ID:          product.ID,
		User:        oUser,
		Brand:       oBrand,
		Category:    oCategory,
		GenreID:     product.GenreID,
		Name:        product.Name,
		Price:       product.Price,
		Image:       product.Image.Filename,
		Description: product.Description,
	}
	oProductInventory := &outputdata.ProductInventory{
		ProductID: product.Inventory.ProductID,
		Receiving: product.Inventory.Receiving,
		Shipping:  product.Inventory.Shipping,
		Disposal:  product.Inventory.Disposal,
		Stock:     product.Inventory.Stock,
	}
	oProductManagement := &outputdata.ProductManagement{
		ProductID: product.Management.ProductID,
		Storage:   product.Management.Storage,
		Memo:      product.Management.Memo,
		Barcode:   product.Management.Barcode,
	}
	return &outputdata.Inventory{
		Product:    oProduct,
		Inventory:  oProductInventory,
		Management: oProductManagement,
	}
}
