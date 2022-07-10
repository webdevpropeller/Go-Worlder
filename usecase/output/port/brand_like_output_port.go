package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// BrandLikeOutputPort ...
type BrandLikeOutputPort interface {
	Index([]model.User) []outputdata.Profile
}
