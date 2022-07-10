package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// BrandLike ...
type BrandLike struct {
	Brand *Brand `validate:"required"`
	User  *User  `validate:"required"`
}

// NewBrandLike ...
func NewBrandLike(brand *Brand, user *User) (*BrandLike, error) {
	brandLike := &BrandLike{
		Brand: brand,
		User:  user,
	}
	err := validator.Struct(brandLike)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return brandLike, nil
}
