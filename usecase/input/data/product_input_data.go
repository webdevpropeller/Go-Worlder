package inputdata

import "mime/multipart"

// Product ...
type Product struct {
	UserID      string `validate:"required"`
	BrandID     string `validate:"required"`
	CategoryID  uint   `validate:"min=1"`
	GenreID     uint
	Name        string `validate:"required"`
	Price       uint   `validate:"min=1"`
	Image       *multipart.FileHeader
	Description string
	IsDraft     bool
}

// NewProduct ...
type NewProduct struct {
	*Product
}

// UpdatedProduct ...
type UpdatedProduct struct {
	ID string
	*Product
}
