package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// AccountNameDatabase ...
type AccountNameDatabase struct {
	SQLHandler
}

// NewAccountNameDatabase ...
func NewAccountNameDatabase(sqlHandler SQLHandler) repository.AccountNameRepository {
	return &AccountNameDatabase{sqlHandler}
}

// FindList ...
func (db *AccountNameDatabase) FindList() ([]model.AccountName, error) {
	accountNameTable := &accountDB.MAccountNamesTable
	accountTypeTable := &accountDB.MAccountTypesTable
	statement := NewSQLBuilder().Select(accountNameTable).
		LeftJoinWithColumns(accountTypeTable, accountTypeTable.ID, accountNameTable.AccountTypeID).
		Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	accountNameList := []model.AccountName{}
	for rows.Next() {
		accountName := model.AccountName{}
		scanList := generateScanList(&accountName)
		err = rows.Scan(scanList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		accountNameList = append(accountNameList, accountName)
	}
	return accountNameList, nil
}

// FindByID ...
func (db *AccountNameDatabase) FindByID(id string) (*model.AccountName, error) {
	accountNameTable := &accountDB.MAccountNamesTable
	accountTypeTable := &accountDB.MAccountTypesTable
	statement := NewSQLBuilder().Select(accountNameTable).
		LeftJoinWithColumns(accountTypeTable, accountTypeTable.ID, accountNameTable.AccountTypeID).
		Match(accountNameTable.ID).
		Statement()
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
	accountName := &model.AccountName{}
	scanList := generateScanList(accountName)
	err = rows.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return accountName, nil
}
