package model

import (
	"go_worlder_system/errs"
	"go_worlder_system/validator"
	"mime/multipart"

	log "github.com/sirupsen/logrus"
)

// Brand is an entity
type Brand struct {
	ID          string    `validate:"required"`
	User        *User     `validate:"required"`
	Category    *Category `validate:"required"`
	Name        string    `validate:"required"`
	Slogan      string
	LogoPath    string
	LogoImage   *multipart.FileHeader
	Description string
	IsDraft     bool
}

// NewBrand ...
func NewBrand(user *User, category *Category, name string, slogan string, logoImage *multipart.FileHeader, description string, isDraft bool) (*Brand, error) {
	id := generateID(idPrefix.BrandID)
	var logoPath string
	if logoImage != nil {
		logoPath = logoImage.Filename
	}
	brand := &Brand{
		ID:          id,
		User:        user,
		Category:    category,
		Name:        name,
		Slogan:      slogan,
		LogoPath:    logoPath,
		LogoImage:   logoImage,
		Description: description,
		IsDraft:     isDraft,
	}
	err := validator.Struct(brand)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return brand, nil
}

// IsOwner ...
func (brand *Brand) IsOwner(userID string) bool {
	return brand.User.ID == userID
}

// Update ...
func (brand *Brand) Update(
	category *Category, name string, slogan string,
	logoImage *multipart.FileHeader, description string, isDraft bool,
) error {
	if brand.IsDraft == false && isDraft == true {
		errMsg := "Brand cannot be returned to draft"
		log.Error(errMsg)
		return errs.Conflict.New(errMsg)
	}
	brand.Category = category
	brand.Name = name
	brand.Slogan = slogan
	brand.LogoImage = logoImage
	brand.Description = description
	brand.IsDraft = isDraft
	err := validator.Struct(brand)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
