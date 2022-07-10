package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// ProjectPresenter ...
type ProjectPresenter struct {
}

// NewProjectPresenter ...
func NewProjectPresenter() outputport.ProjectOutputPort {
	return &ProjectPresenter{}
}

// Index ...
func (presenter *ProjectPresenter) Index(projectList []model.Project) []outputdata.Project {
	oProjectList := []outputdata.Project{}
	for _, project := range projectList {
		oProject := presenter.convert(&project)
		oProjectList = append(oProjectList, *oProject)
	}
	return oProjectList
}

// Show ...
func (presenter *ProjectPresenter) Show(project *model.Project) *outputdata.Project {
	return presenter.convert(project)
}

// Edit ...
func (presenter *ProjectPresenter) Edit(project *model.Project) *outputdata.Project {
	return presenter.convert(project)
}

func (presenter *ProjectPresenter) convert(project *model.Project) *outputdata.Project {
	oUser := &outputdata.UserSimplified{
		ID:   project.User.ID,
		Name: project.User.Profile.Company,
	}
	return &outputdata.Project{
		ID:   project.ID,
		User: oUser,
		Name: project.Name,
	}
}
