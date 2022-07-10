package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProfilePresenter ...
type ProfilePresenter struct {
}

// NewProfilePresenter ...
func NewProfilePresenter() *ProfilePresenter {
	return &ProfilePresenter{}
}

// Show ...
func (presenter *ProfilePresenter) Show(profile *model.Profile) *outputdata.PublicUser {
	return &outputdata.PublicUser{
		ID:   profile.UserID,
		Name: profile.Company,
	}
}

func (presenter *ProfilePresenter) Create(user *model.User) *outputdata.User {
	oProfile := presenter.convert(user.Profile)
	return &outputdata.User{
		ID:      user.ID,
		Email:   user.Email,
		Profile: oProfile,
	}
}

func (presenter *ProfilePresenter) New(industries model.Industries, countries model.Countries, cardCompanies model.CardCompanies) *outputdata.ProfileSelectItem {
	oIndustries := convertOption(industries)
	oCountries := convertOption(countries)
	oCardCompanies := convertOption(cardCompanies)
	return &outputdata.ProfileSelectItem{
		Industries:    oIndustries,
		Countries:     oCountries,
		CardCompanies: oCardCompanies,
	}
}

// Edit ...
func (presenter *ProfilePresenter) Edit(profile *model.Profile) *outputdata.Profile {
	return presenter.convert(profile)
}

// Profile ...
func (presenter *ProfilePresenter) Profile(profile *model.Profile, brandList []model.Brand, productList []model.Product, projectList []model.Project) *outputdata.Profile {
	oProfile := presenter.convert(profile)
	oBrandList := []outputdata.Brand{}
	for _, brand := range brandList {
		oBrand := outputdata.Brand{
			ID:          brand.ID,
			Name:        brand.Name,
			Slogan:      brand.Slogan,
			LogoImage:   brand.LogoImage.Filename,
			Description: brand.Description,
		}
		oBrandList = append(oBrandList, oBrand)
	}
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
		oProduct := outputdata.Product{
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
		oProductList = append(oProductList, oProduct)
	}
	oProjectList := []outputdata.Project{}
	for _, project := range projectList {
		oProject := outputdata.Project{
			ID:   project.ID,
			Name: project.Name,
		}
		oProjectList = append(oProjectList, oProject)
	}
	return oProfile

}

// BrandLike ...
func (presenter *ProfilePresenter) BrandLike(brandList []model.Brand) []outputdata.Brand {
	return []outputdata.Brand{}
}

func (presenter *ProfilePresenter) convert(profile *model.Profile) *outputdata.Profile {
	if profile == nil {
		return nil
	}
	return &outputdata.Profile{
		Activity:   profile.Activity,
		Industry:   profile.Industry,
		Company:    profile.Company,
		Country:    profile.Country,
		Address1:   profile.Address1,
		Address2:   profile.Address2,
		URL:        profile.URL,
		Phone:      profile.Phone,
		Logo:       profile.LogoPath,
		FirstName:  profile.FirstName,
		MiddleName: profile.MiddleName,
		FamilyName: profile.FamilyName,
		Icon:       profile.IconPath,
	}
}
