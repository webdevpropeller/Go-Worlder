package repository

import "go_worlder_system/domain/model"

// ProjectRepository ...
type ProjectRepository interface {
	FindListByUserID(userID string) ([]model.Project, error)
	FindByID(id string) (*model.Project, error)
	Save(*model.Project) error
	Update(*model.Project) error
	Delete(*model.Project) error
}
