package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// BrandPresenter ...
type BrandPresenter struct {
}

// NewBrandPresenter ...
func NewBrandPresenter() outputport.BrandOutputPort {
	return &BrandPresenter{}
}

// Index ...
func (presenter *BrandPresenter) Index(brandList []model.Brand) []outputdata.Brand {
	oBrandList := []outputdata.Brand{}
	for _, brand := range brandList {
		oBrand := presenter.convert(&brand)
		oBrandList = append(oBrandList, *oBrand)
	}
	return oBrandList
}

// Show ...
func (presenter *BrandPresenter) Show(brand *model.Brand) *outputdata.Brand {
	return presenter.convert(brand)
}

// Edit ...
func (presenter *BrandPresenter) Edit(brand *model.Brand) *outputdata.Brand {
	return presenter.convert(brand)
}

func (presenter *BrandPresenter) convert(brand *model.Brand) *outputdata.Brand {
	oUser := &outputdata.UserSimplified{
		ID:   brand.User.ID,
		Name: brand.User.Profile.Company,
	}
	oCategory := &outputdata.Category{
		ID:   brand.Category.ID,
		Name: brand.Category.Name,
	}
	return &outputdata.Brand{
		ID:          brand.ID,
		User:        oUser,
		Category:    oCategory,
		Name:        brand.Name,
		Slogan:      brand.Slogan,
		LogoImage:   brand.LogoPath,
		Description: brand.Description,
	}
}
