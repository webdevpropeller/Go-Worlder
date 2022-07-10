package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// ProductPresenter ...
type ProductPresenter struct {
}

// NewProductPresenter ...
func NewProductPresenter() outputport.ProductOutputPort {
	return &ProductPresenter{}
}

// Index ...
func (presenter *ProductPresenter) Index(productList []model.Product) []outputdata.Product {
	oProductList := []outputdata.Product{}
	for _, product := range productList {
		oProduct := presenter.convert(&product)
		oProductList = append(oProductList, *oProduct)
	}
	return oProductList
}

// Show ...
func (presenter *ProductPresenter) Show(product *model.Product) *outputdata.Product {
	return presenter.convert(product)
}

// Edit ...
func (presenter *ProductPresenter) Edit(product *model.Product) *outputdata.Product {
	return presenter.convert(product)
}

func (presenter *ProductPresenter) convert(product *model.Product) *outputdata.Product {
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
	return &outputdata.Product{
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
}
