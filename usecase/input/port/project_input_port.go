package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProjectInputPort ...
type ProjectInputPort interface {
	Index(userID string) ([]outputdata.Project, error)
	Show(id string) (*outputdata.Project, error)
	Edit(id string, userID string) (*outputdata.Project, error)
	Create(*inputdata.NewProject) error
	Update(*inputdata.UpdatedProject) error
	Delete(id string, userID string) error
}
