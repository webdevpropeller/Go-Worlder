package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// AccountTypeDatabase ...
type AccountTypeDatabase struct {
	SQLHandler
}

// NewAccountTypeDatabase ...
func NewAccountTypeDatabase(sqlHandler SQLHandler) repository.AccountTypeRepository {
	return &AccountTypeDatabase{sqlHandler}
}

// FindByID ...
func (db AccountTypeDatabase) FindByID(id uint) (*model.AccountType, error) {
	mAccountTypesTable := &accountDB.MAccountTypesTable
	statement := NewSQLBuilder().Select(mAccountTypesTable).Match(mAccountTypesTable.ID).Statement()
	valueList := []interface{}{id}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	if !rows.Next() {
		errMsg := "The account name doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	accountType := &model.AccountType{}
	scanList := generateScanList(accountType)
	err = rows.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return accountType, nil
}
