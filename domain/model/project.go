package model

import (
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// Project is an entity
type Project struct {
	ID   string
	User *User
	Name string
}

// NewProject ...
func NewProject(user *User, name string) (*Project, error) {
	id := generateID(idPrefix.ProjectID)
	project := &Project{
		ID:   id,
		User: user,
		Name: name,
	}
	err := validator.Struct(project)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return project, nil
}

// IsOwner ...
func (project *Project) IsOwner(userID string) bool {
	return project.User.ID == userID
}

// Update ...
func (project *Project) Update(name string) error {
	project.Name = name
	err := validator.Struct(project)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
