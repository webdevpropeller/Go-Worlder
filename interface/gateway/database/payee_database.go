package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// PayeeDatabase ...
type PayeeDatabase struct {
	SQLHandler
}

// NewPayeeDatabase ...
func NewPayeeDatabase(sqlHandler SQLHandler) repository.PayeeRepository {
	return &PayeeDatabase{sqlHandler}
}

// FindListByUserID ...
func (db *PayeeDatabase) FindListByUserID(userID string) ([]model.Payee, error) {
	payeesTable := &transactionDB.PayeesTable
	statement := db.findStatement(payeesTable.UserID)
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	payeeList := []model.Payee{}
	for rows.Next() {
		payee := model.Payee{}
		scanList := generateScanList(&payee)
		rows.Scan(scanList...)
		payeeList = append(payeeList, payee)
	}
	return payeeList, nil
}

// FindByID ...
func (db *PayeeDatabase) FindByID(id string) (*model.Payee, error) {
	payeesTable := &transactionDB.PayeesTable
	statement := db.findStatement(payeesTable.ID)
	valueList := []interface{}{id}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The account doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	payee := &model.Payee{}
	scanList := generateScanList(payee)
	err = row.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return payee, nil
}

// Save ...
func (db *PayeeDatabase) Save(payee *model.Payee) error {
	_, err := db.Transaction(func() (interface{}, error) {
		transactionPartnersTable := &transactionDB.TransactionPartnersTable
		statement := NewSQLBuilder().Insert(transactionPartnersTable)
		valueList := []interface{}{payee.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		payeesTable := &transactionDB.PayeesTable
		statement = NewSQLBuilder().Insert(payeesTable)
		valueList = generateValueList(payee)
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
func (db *PayeeDatabase) Update(payee *model.Payee) error {
	_, err := db.Transaction(func() (interface{}, error) {
		payeesTable := &transactionDB.PayeesTable
		statement := NewSQLBuilder().Update(payeesTable)
		valueList := generateValueList(payee)
		valueList = append(valueList, payee.ID)
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
func (db *PayeeDatabase) Delete(payee *model.Payee) error {
	_, err := db.Transaction(func() (interface{}, error) {
		payeesTable := &transactionDB.PayeesTable
		statement := NewSQLBuilder().Delete(payeesTable).Match(payeesTable.ID).Statement()
		valueList := []interface{}{payee.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// findStatement ...
func (db *PayeeDatabase) findStatement(condition string) string {
	payeesTable := &transactionDB.PayeesTable
	userInfoTable := &userDB.ProfileTable
	statement := NewSQLBuilder().Select(payeesTable).
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, payeesTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		Match(condition).
		Statement()
	return statement
}
