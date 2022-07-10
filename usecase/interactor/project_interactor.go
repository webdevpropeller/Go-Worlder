package interactor

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// ProjectInteractor ...
type ProjectInteractor struct {
	outputport        outputport.ProjectOutputPort
	userRepository    repository.UserRepository
	projectRepository repository.ProjectRepository
}

// NewProjectInteractor ...
func NewProjectInteractor(
	outputport outputport.ProjectOutputPort,
	userRepository repository.UserRepository,
	projectRepository repository.ProjectRepository,
) *ProjectInteractor {
	return &ProjectInteractor{
		outputport:        outputport,
		userRepository:    userRepository,
		projectRepository: projectRepository,
	}
}

// Index ...
func (interactor *ProjectInteractor) Index(userID string) ([]outputdata.Project, error) {
	projectList, err := interactor.projectRepository.FindListByUserID(userID)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return interactor.outputport.Index(projectList), nil
}

// Show ...
func (interactor *ProjectInteractor) Show(id string) (*outputdata.Project, error) {
	project, err := interactor.projectRepository.FindByID(id)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return interactor.outputport.Show(project), nil
}

// Edit ...
func (interactor *ProjectInteractor) Edit(id string, userID string) (*outputdata.Project, error) {
	project, err := interactor.projectRepository.FindByID(id)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	if !project.IsOwner(userID) {
		errMsg := "The user can't get the brand"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return interactor.outputport.Edit(project), nil
}

// Create ...
func (interactor *ProjectInteractor) Create(iNewProject *inputdata.NewProject) error {
	user, err := interactor.userRepository.FindByID(iNewProject.UserID)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil
	}
	project, err := model.NewProject(user, iNewProject.Name)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	err = interactor.projectRepository.Save(project)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

// Update ...
func (interactor *ProjectInteractor) Update(iUpdatedProject *inputdata.UpdatedProject) error {
	project, err := interactor.projectRepository.FindByID(iUpdatedProject.ID)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	if !project.IsOwner(iUpdatedProject.UserID) {
		errMsg := "The user can't update the project"
		log.Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = project.Update(iUpdatedProject.Name)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	err = interactor.projectRepository.Update(project)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *ProjectInteractor) Delete(id string, userID string) error {
	project, err := interactor.projectRepository.FindByID(id)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	if !project.IsOwner(userID) {
		errMsg := "The user can't delete the project"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error()
		return errs.Forbidden.New(errMsg)
	}
	err = interactor.projectRepository.Delete(project)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return err
	}
	return nil
}
