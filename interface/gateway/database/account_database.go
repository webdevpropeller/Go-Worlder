package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"
	"strings"

	log "github.com/sirupsen/logrus"
)

// AccountDatabase ...
type AccountDatabase struct {
	SQLHandler
}

// NewAccountDatabase ...
func NewAccountDatabase(sqlHandler SQLHandler) repository.AccountRepository {
	return &AccountDatabase{sqlHandler}
}

// FindListByUserID ...
func (db *AccountDatabase) FindListByUserID(userID string) ([]model.Account, error) {
	accountsTable := &accountDB.AccountsTable
	statement := db.findStatement(accountsTable.UserID)
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	accountList := []model.Account{}
	for rows.Next() {
		account := model.Account{}
		scanList := generateScanList(&account)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		accountList = append(accountList, account)
	}
	return accountList, nil
}

// FindByID ...
func (db *AccountDatabase) FindByID(id string) (*model.Account, error) {
	accountsTable := &accountDB.AccountsTable
	statement := db.findStatement(accountsTable.ID)
	row, err := db.Query(statement, id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exist := row.Next()
	if !exist {
		errMsg := "The account doesn't exist"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	account := &model.Account{}
	scanList := generateScanList(account)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return account, nil
}

// FindCashByUserID ...
func (db *AccountDatabase) FindCashByUserID(userID string) (*model.Cash, error) {
	// TODO
	transactionsTable := &transactionDB.TransactionsTable
	statement := strings.Join([]string{
		rw.Select,
		"SUM(",
		rw.Case, transactionsTable.IsIncome,
		rw.When, trueVal, rw.Then, transactionsTable.Amount,
		rw.When, falseVal, rw.Then, transactionsTable.Amount, "* -1",
		rw.Else, "0",
		rw.End,
		")",
		rw.Where, transactionsTable.UserID, "= ?",
	}, " ")
	row, err := db.Query(statement, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exist := row.Next()
	var cash model.Cash = 0
	if !exist {
		return &cash, nil
	}
	err = row.Scan(&cash)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return &cash, nil
}

// Save ...
func (db *AccountDatabase) Save(account *model.Account) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Insert into transaction account table
		transactionAccountTable := &transactionDB.TransactionAccountsTable
		statement := NewSQLBuilder().Insert(transactionAccountTable)
		_, err := db.Exec(statement, account.ID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into account table
		accountsTable := &accountDB.AccountsTable
		statement = NewSQLBuilder().Insert(accountsTable)
		valueList := generateValueList(account)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into active table
		accountsActiveTable := &accountDB.AccountsActiveTable
		statement = NewSQLBuilder().Insert(accountsActiveTable)
		valueList = []interface{}{account.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *AccountDatabase) Delete(account *model.Account) error {
	_, err := db.Transaction(func() (interface{}, error) {
		accountsActiveTable := &accountDB.AccountsActiveTable
		valueList := []interface{}{account.ID}
		// Delete from active table
		statement := NewSQLBuilder().Delete(accountsActiveTable).Match(accountsActiveTable.AccountID).Statement()
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into deleted table
		accountsDeletedTable := &accountDB.AccountsDeletedTable
		statement = NewSQLBuilder().Insert(accountsDeletedTable)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (db *AccountDatabase) findStatement(condition string) string {
	accountsTable := &accountDB.AccountsTable
	accountsActiveTable := &accountDB.AccountsActiveTable
	userInfoTable := &userDB.ProfileTable
	mAccountNameTable := &accountDB.MAccountNamesTable
	mAccountTypeTable := &accountDB.MAccountTypesTable
	transactionsTable := &transactionDB.TransactionsTable
	statement := NewSQLBuilder().Select(accountsTable).
		RightJoin(accountsActiveTable, accountsActiveTable.AccountID, accountsTable.ID).
		// user
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, accountsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		// account name
		LeftJoinWithColumns(mAccountNameTable, mAccountNameTable.ID, accountsTable.AccountNameID).
		LeftJoinWithColumns(mAccountTypeTable, mAccountTypeTable.ID, mAccountNameTable.AccountTypeID).
		// account balance
		LeftJoin(transactionsTable, transactionsTable.TransactionAccountID, accountsTable.ID).
		Column(strings.Join([]string{
			"SUM(",
			rw.Case, transactionsTable.IsIncome,
			rw.When, trueVal, rw.Then, transactionsTable.Amount,
			rw.When, falseVal, rw.Then, transactionsTable.Amount, "* -1",
			rw.Else, "0",
			rw.End,
			")",
		}, " ")).
		Match(condition).
		Statement()
	return statement
}
