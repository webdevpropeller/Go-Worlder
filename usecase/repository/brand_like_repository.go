package repository

import "go_worlder_system/domain/model"

// BrandLikeRepository ...
type BrandLikeRepository interface {
	FindByBrandIDAndUserID(brandID string, userID string) (*model.BrandLike, error)
	Save(*model.BrandLike) error
	Delete(*model.BrandLike) error
}
