package repository

import (
	"go_worlder_system/domain/model"
)

// UserRepository is an interface for DB operation
type UserRepository interface {
	FindByID(id string) (*model.User, error)
	FindByEmail(email string) (*model.User, error)
	FindByAccountID(string) (*model.Profile, error)
	FindCardByUserID(userID string) (*model.Card, error)
	FindListByLikeBrandID(brandID string) ([]model.User, error)
	Save(*model.User) error
	SaveProfile(*model.User) error
	SaveCard(*model.Card) error
	DeleteCard(*model.Card) error
	UpdateToActive(*model.User) error
	UpdatePassword(*model.User) error
	UpdateProfile(*model.User) error
}
