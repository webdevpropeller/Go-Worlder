package service

import "go_worlder_system/domain/model"

// BrandProduct ...
type BrandProduct struct {
	Brand       *model.Brand
	ProductList []model.Product
}

// NewBrandProduct ...
func NewBrandProduct(brand *model.Brand, productList []model.Product) *BrandProduct {
	return &BrandProduct{
		Brand:       brand,
		ProductList: productList,
	}
}
