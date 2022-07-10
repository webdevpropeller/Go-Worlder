package inputport

import outputdata "go_worlder_system/usecase/output/data"

// BrandLikeInputPort ...
type BrandLikeInputPort interface {
	Index(id string) ([]outputdata.Profile, error)
	Create(id string, userID string) error
	Delete(id string, userID string) error
}
