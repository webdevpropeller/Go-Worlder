package datasource

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/interface/gateway/database"
)

// UserDataSource ...
type UserDataSource struct {
	database *database.UserDatabase
}

// NewUserDataSource ...
func NewUserDataSource(database *database.UserDatabase) *UserDataSource {
	return &UserDataSource{
		database: database,
	}
}

// FindByID ...
func (dc *UserDataSource) FindByID(id string) (*model.User, error) {
	return dc.database.FindByID(id)
}

// FindByEmail ...
func (dc *UserDataSource) FindByEmail(email string) (*model.User, error) {
	return dc.database.FindByEmail(email)
}

// FindListByLikeBrandID ...
func (dc *UserDataSource) FindListByLikeBrandID(brandID string) ([]model.User, error) {
	return dc.database.FindListByLikeBrandID(brandID)
}

// Save ...
func (dc *UserDataSource) Save(user *model.User) error {
	return dc.database.Save(user)
}

// SaveInfo ...
func (dc *UserDataSource) SaveProfile(user *model.User) error {
	return dc.database.SaveProfile(user)
}

// UpdateToActive ...
func (dc *UserDataSource) UpdateToActive(user *model.User) error {
	return dc.database.UpdateToActive(user)
}

// UpdatePassword ...
func (dc *UserDataSource) UpdatePassword(user *model.User) error {
	return dc.database.UpdatePassword(user)
}

// UpdateInfo ...
func (dc *UserDataSource) UpdateProfile(user *model.User) error {
	return dc.database.UpdateProfile(user)
}
