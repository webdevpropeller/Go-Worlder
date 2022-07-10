package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// CategoryDatabase ...
type CategoryDatabase struct {
	SQLHandler
}

// NewCategoryDatabase ...
func NewCategoryDatabase(sqlHandler SQLHandler) repository.CategoryRepository {
	return &CategoryDatabase{sqlHandler}
}

// FindByID ...
func (db *CategoryDatabase) FindByID(id uint) (*model.Category, error) {
	mCategoriesTable := &categoryDB.MCategoriesTable
	statement := NewSQLBuilder().Select(mCategoriesTable).Match(mCategoriesTable.ID).Statement()
	valueList := []interface{}{id}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	exists := rows.Next()
	if !exists {
		errMsg := "The category doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	category := &model.Category{}
	scanList := generateScanList(category)
	err = rows.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return category, nil
}
