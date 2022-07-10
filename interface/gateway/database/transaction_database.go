package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"
	"strings"

	log "github.com/sirupsen/logrus"
)

// TransactionDatabase ...
type TransactionDatabase struct {
	SQLHandler
}

// NewTransactionDatabase ...
func NewTransactionDatabase(sqlHandler SQLHandler) repository.TransactionRepository {
	return &TransactionDatabase{sqlHandler}
}

// FindListByUserID ...
func (db *TransactionDatabase) FindListByUserID(userID string) ([]model.Transaction, error) {
	transactionsTable := &transactionDB.TransactionsTable
	statement := db.findStatement(transactionsTable.UserID)
	rows, err := db.Query(statement, userID)
	if err != nil {
		log.WithFields(log.Fields{"statement": statement}).Error(err)
		return nil, err
	}
	defer rows.Close()
	transactionList := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		scanList := generateScanList(&transaction)
		err = rows.Scan(scanList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			break
		}
		transactionList = append(transactionList, transaction)
	}
	return transactionList, nil
}

// FindListByAccountID ...
func (db *TransactionDatabase) FindListByAccountID(accountID string) ([]model.Transaction, error) {
	transactionAccountsTable := &transactionDB.TransactionAccountsTable
	statement := db.findStatement(transactionAccountsTable.ID)
	rows, err := db.Query(statement, accountID)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	transactionList := []model.Transaction{}
	for rows.Next() {
		transaction := model.Transaction{}
		scanList := generateScanList(&transaction)
		err = rows.Scan(scanList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			break
		}
		transactionList = append(transactionList, transaction)
	}
	return transactionList, nil
}

// FindByID ...
func (db *TransactionDatabase) FindByID(id string) (*model.Transaction, error) {
	transactionsTable := &transactionDB.TransactionsTable
	statement := db.findStatement(transactionsTable.ID)
	valueList := []interface{}{id}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	exists := rows.Next()
	if !exists {
		errMsg := "The transaction is not found"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	transaction := &model.Transaction{}
	scanList := generateScanList(transaction)
	err = rows.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return transaction, nil
}

// FindAccountItemList ...
func (db *TransactionDatabase) FindAccountItemList() ([]model.AccountItem, error) {
	mAccountItemsTable := &transactionDB.MAccountItemsTable
	statement := NewSQLBuilder().Select(mAccountItemsTable).Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	accountItemList := []model.AccountItem{}
	for rows.Next() {
		accountItem := model.AccountItem{}
		scanList := generateScanList(&accountItem)
		rows.Scan(scanList...)
		accountItemList = append(accountItemList, accountItem)
	}
	return accountItemList, nil
}

// FindAccountItemByID ...
func (db *TransactionDatabase) FindAccountItemByID(id string) (*model.AccountItem, error) {
	mAccountItemsTable := &transactionDB.MAccountItemsTable
	statement := NewSQLBuilder().Select(mAccountItemsTable).Match(mAccountItemsTable.ID).Statement()
	row, err := db.Query(statement, id)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "Not found account item"
		log.WithFields(log.Fields{}).Error(err)
		return nil, errs.NotFound.New(errMsg)
	}
	accountItem := &model.AccountItem{}
	scanList := generateScanList(accountItem)
	row.Scan(scanList...)
	return accountItem, nil
}

// FindPartnerList ...
func (db *TransactionDatabase) FindPartnerList() ([]model.TransactionPartner, error) {
	statement := NewSQLBuilder().Select(customeTransactionPartnerTable).Statement()
	valueList := []interface{}{partnerID}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	partnerList := []model.TransactionPartner{}
	for row.Next() {
		partner := model.TransactionPartner{}
		scanList := generateScanList(&partner)
		err = row.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		partnerList = append(partnerList, partner)
	}
	return partnerList, nil
}

// FindPartnerByID ...
func (db *TransactionDatabase) FindPartnerByID(partnerID string) (*model.TransactionPartner, error) {
	statement := NewSQLBuilder().Select(customeTransactionPartnerTable).Match(customeTransactionPartnerTable.ID).Statement()
	valueList := []interface{}{partnerID}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The partner is not found"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	partner := &model.TransactionPartner{}
	scanList := generateScanList(partner)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return partner, nil
}

// FindAccountByID ...
func (db *TransactionDatabase) FindAccountByID(accountID string) (*model.TransactionAccount, error) {
	if accountID == "cash" {
		return &model.TransactionAccount{ID: accountID, Name: "cash"}, nil
	}
	accountTable := &accountDB.AccountsTable
	maccountNameTable := &accountDB.MAccountNamesTable
	columnList := []string{accountTable.ID, maccountNameTable.Name}
	columnQuery := generateSelectColumnQuery(columnList)
	statement := strings.Join([]string{
		rw.Select, columnQuery,
		rw.From, accountTable.NAME(),
		rw.LeftJoin, maccountNameTable.NAME(),
		rw.On, accountTable.AccountNameID, "=", maccountNameTable.ID,
		rw.Where, accountTable.ID, "= ?",
	}, " ")
	row, err := db.Query(statement, accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The partner is not found"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	account := &model.TransactionAccount{}
	scanList := generateScanList(account)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return account, nil
}

// Save ...
func (db *TransactionDatabase) Save(transaction *model.Transaction) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Insert into transactions
		transactionsTable := &transactionDB.TransactionsTable
		statement := NewSQLBuilder().Insert(transactionsTable)
		valueList := generateValueList(transaction)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Update ...
func (db *TransactionDatabase) Update(transaction *model.Transaction) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Update transactions
		transactionsTable := &transactionDB.TransactionsTable
		statement := NewSQLBuilder().Update(transactionsTable)
		valueList := generateValueList(transaction)
		valueList = append(valueList, transaction.ID)
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
func (db *TransactionDatabase) Delete(transaction *model.Transaction) error {
	_, err := db.Transaction(func() (interface{}, error) {
		transactionsTable := &transactionDB.TransactionsTable
		statement := NewSQLBuilder().Delete(transactionsTable).Match(transactionsTable.ID).Statement()
		_, err := db.Exec(statement, transaction.ID)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (db *TransactionDatabase) findStatement(condition string) string {
	transactionsTable := &transactionDB.TransactionsTable
	userInfoTable := &userDB.ProfileTable
	maccountItemsTable := &transactionDB.MAccountItemsTable
	statement := NewSQLBuilder().Select(transactionsTable).
		// user
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, transactionsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		// account item
		LeftJoinWithColumns(maccountItemsTable, maccountItemsTable.ID, transactionsTable.AccountItemID).
		// transaction partner
		LeftJoinWithColumns(customeTransactionPartnerTable, customeTransactionPartnerTable.ID, transactionsTable.TransactionPartnerID).
		// account
		LeftJoinWithColumns(customeTransactionAccountTable, customeTransactionAccountTable.ID, transactionsTable.TransactionAccountID).
		Match(condition).
		Statement()
	return statement
}
