package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// BrandLikePresenter ...
type BrandLikePresenter struct {
}

// NewBrandLikePresenter ...
func NewBrandLikePresenter() outputport.BrandLikeOutputPort {
	return &BrandLikePresenter{}
}

// Index ...
func (presenter *BrandLikePresenter) Index(userList []model.User) []outputdata.Profile {
	oProfileList := []outputdata.Profile{}
	for _, user := range userList {
		oProfile := outputdata.Profile{
			Activity:   user.Profile.Activity,
			Industry:   user.Profile.Industry,
			Company:    user.Profile.Company,
			Country:    user.Profile.Country,
			Address1:   user.Profile.Address1,
			Address2:   user.Profile.Address2,
			URL:        user.Profile.URL,
			Phone:      user.Profile.Phone,
			Logo:       user.Profile.Logo.Filename,
			FirstName:  user.Profile.FirstName,
			MiddleName: user.Profile.MiddleName,
			FamilyName: user.Profile.FamilyName,
			Icon:       user.Profile.Icon.Filename,
		}
		oProfileList = append(oProfileList, oProfile)
	}
	return oProfileList
}
