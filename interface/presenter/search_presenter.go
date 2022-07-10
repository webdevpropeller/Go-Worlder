package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// SearchPresenter ...
type SearchPresenter struct {
}

// NewSearchPresenter ...
func NewSearchPresenter() outputport.SearchOutputPort {
	return &SearchPresenter{}
}

// SearchProduct ...
func (presenter *SearchPresenter) SearchProduct(productList []model.Product) []outputdata.Product {
	oProductList := []outputdata.Product{}
	for _, product := range productList {
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
		oProductList = append(oProductList, *oProduct)
	}
	return oProductList
}
