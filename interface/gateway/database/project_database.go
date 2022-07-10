package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// ProjectDatabase ...
type ProjectDatabase struct {
	SQLHandler
}

// NewProjectDatabase ...
func NewProjectDatabase(sqlHandler SQLHandler) repository.ProjectRepository {
	return &ProjectDatabase{sqlHandler}
}

// FindListByUserID ...
func (db *ProjectDatabase) FindListByUserID(userID string) ([]model.Project, error) {
	projectsTable := &projectDB.ProjectsTable
	statement := db.findStatement(projectsTable.UserID)
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	projectList := []model.Project{}
	for rows.Next() {
		project := model.Project{}
		scanList := generateScanList(&project)
		rows.Scan(scanList...)
		projectList = append(projectList, project)
	}
	return projectList, nil
}

// FindByID ...
func (db *ProjectDatabase) FindByID(id string) (*model.Project, error) {
	projectsTable := &projectDB.ProjectsTable
	statement := db.findStatement(projectsTable.ID)
	valueList := []interface{}{id}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The project doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	project := &model.Project{}
	scanList := generateScanList(project)
	err = row.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return project, nil
}

// Save ...
func (db *ProjectDatabase) Save(project *model.Project) error {
	_, err := db.Transaction(func() (interface{}, error) {
		projectsTable := &projectDB.ProjectsTable
		statement := NewSQLBuilder().Insert(projectsTable)
		valueList := generateValueList(project)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		// Insert into active table
		projectsActiveTable := &projectDB.ProjectsActiveTable
		statement = NewSQLBuilder().Insert(projectsActiveTable)
		valueList = []interface{}{project.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Update ...
func (db *ProjectDatabase) Update(project *model.Project) error {
	_, err := db.Transaction(func() (interface{}, error) {
		projectsTable := &projectDB.ProjectsTable
		statement := NewSQLBuilder().Update(projectsTable)
		valueList := generateValueList(project)
		valueList = append(valueList, project.ID)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *ProjectDatabase) Delete(project *model.Project) (err error) {
	_, err = db.Transaction(func() (interface{}, error) {
		projectsActiveTable := &projectDB.ProjectsActiveTable
		statement := NewSQLBuilder().Delete(projectsActiveTable).Match(projectsActiveTable.ProjectID).Statement()
		valueList := []interface{}{project.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		projectsDeletedTable := &projectDB.ProjectsDeletedTable
		statement = NewSQLBuilder().Insert(projectsDeletedTable)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return
}

func (db *ProjectDatabase) findStatement(condition string) string {
	projectsTable := &projectDB.ProjectsTable
	projectsActiveTable := &projectDB.ProjectsActiveTable
	userInfoTable := &userDB.ProfileTable
	statement := NewSQLBuilder().Select(projectsTable).
		RightJoin(projectsActiveTable, projectsActiveTable.ProjectID, projectsTable.ID).
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, projectsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		Match(condition).
		Statement()
	return statement
}
