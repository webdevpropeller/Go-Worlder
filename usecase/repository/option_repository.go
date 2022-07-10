package repository

import "go_worlder_system/domain/model"

type OptionRespository interface {
	FindCountries() (model.Countries, error)
	FindIndustries() (model.Industries, error)
	FindCardCompanies() (model.CardCompanies, error)
	FindCategories() (model.Categories, error)
}
