package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProfileOutputPort ...
type ProfileOutputPort interface {
	Show(*model.Profile) *outputdata.PublicUser
	Create(*model.User) *outputdata.User
	New(model.Industries, model.Countries, model.CardCompanies) *outputdata.ProfileSelectItem
	Edit(*model.Profile) *outputdata.Profile
	Profile(*model.Profile, []model.Brand, []model.Product, []model.Project) *outputdata.Profile
	BrandLike([]model.Brand) []outputdata.Brand
}
