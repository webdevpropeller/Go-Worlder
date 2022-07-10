package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProjectOutputPort ...
type ProjectOutputPort interface {
	Index([]model.Project) []outputdata.Project
	Show(*model.Project) *outputdata.Project
	Edit(*model.Project) *outputdata.Project
}
