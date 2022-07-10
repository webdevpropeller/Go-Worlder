package inputdata

import "mime/multipart"

// Brand ...
type Brand struct {
	UserID      string `validate:"required"`
	CategoryID  uint   `validate:"required"`
	Name        string `validate:"required"`
	Slogan      string
	LogoImage   *multipart.FileHeader
	Description string
	IsDraft     bool
}

// NewBrand ...
type NewBrand struct {
	*Brand
}

// UpdatedBrand ...
type UpdatedBrand struct {
	ID string
	*Brand
}
