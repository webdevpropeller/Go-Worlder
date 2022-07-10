package database

import "strings"

// CustomeUserTable ...
type CustomeUserTable struct {
	table
	UserID         string
	Email          string
	PasswordDigest string
	Category       string
}

// NewCustomeUserTable ...
func NewCustomeUserTable() *CustomeUserTable {
	userAuthTable := &userDB.UserAuthTable
	brandOwnersTable := &userDB.BrandOwnersTable
	partnersTable := &userDB.PartnersTable
	tableName := strings.Join([]string{
		userAuthTable.NAME(),
		rw.LeftJoin, "(",
		rw.Select,
		brandOwnersTable.UserID, ",",
		"1", rw.As, "category",
		rw.From, brandOwnersTable.NAME(),
		rw.UnionAll,
		rw.Select,
		partnersTable.UserID, ",",
		"2", rw.As, "category",
		rw.From, partnersTable.NAME(),
		")", rw.As, "category",
		rw.On, userAuthTable.UserID, "=", "category.user_id",
	}, " ")
	table := table{
		name: tableName,
	}
	return &CustomeUserTable{
		table:          table,
		UserID:         "user_auth.user_id",
		Email:          "user_auth.email",
		PasswordDigest: "user_auth.password_digest",
		Category:       "category.category",
	}
}

// CustomeTransactionPartnerTable ...
type CustomeTransactionPartnerTable struct {
	table
	ID   string
	Name string
	Type string
}

// NewCustomeTransactionPartnerTable ...
func NewCustomeTransactionPartnerTable() *CustomeTransactionPartnerTable {
	userInfoTable := &userDB.ProfileTable
	payeesTable := &transactionDB.PayeesTable
	IDColumn := "id"
	NameColumn := "name"
	typeColumn := "type"
	transactionPartner := "transaction_partner"
	tableName := strings.Join([]string{
		"(",
		rw.Select,
		userInfoTable.UserID, rw.As, IDColumn, ",",
		userInfoTable.Company, rw.As, NameColumn, ",",
		"\"vendor\"", rw.As, typeColumn,
		rw.From, userInfoTable.NAME(),
		rw.UnionAll,
		rw.Select,
		payeesTable.ID, rw.As, IDColumn, ",",
		payeesTable.Name, rw.As, NameColumn, ",",
		"\"payee\"", rw.As, typeColumn,
		rw.From, payeesTable.NAME(),
		")", rw.As, transactionPartner,
	}, " ")
	table := table{
		name: tableName,
	}
	return &CustomeTransactionPartnerTable{
		table: table,
		ID:    strings.Join([]string{transactionPartner, IDColumn}, "."),
		Name:  strings.Join([]string{transactionPartner, NameColumn}, "."),
		Type:  strings.Join([]string{transactionPartner, typeColumn}, "."),
	}
}

// CustomeTransactionAccountTable ...
type CustomeTransactionAccountTable struct {
	table
	ID   string
	Name string
}

// NewCustomeTransactionAccountTable ...
func NewCustomeTransactionAccountTable() *CustomeTransactionAccountTable {
	transactionAccountsTable := &transactionDB.TransactionAccountsTable
	accountsTable := &accountDB.AccountsTable
	maccountNamesTalbe := &accountDB.MAccountNamesTable
	transactionAccounts := "transaction_accounts"
	idColumn := "id"
	nameColumn := "name"
	tableName := strings.Join([]string{
		"(",
		rw.Select,
		transactionAccountsTable.ID, ",",
		rw.Ifnull, "(", maccountNamesTalbe.Name, ",", "\"cash\"", ")", rw.As, nameColumn,
		rw.From, transactionAccountsTable.NAME(),
		rw.LeftJoin, accountsTable.NAME(),
		rw.On, accountsTable.ID, "=", transactionAccountsTable.ID,
		rw.LeftJoin, maccountNamesTalbe.NAME(),
		rw.On, maccountNamesTalbe.ID, "=", accountsTable.AccountNameID,
		")", rw.As, transactionAccounts,
	}, " ")
	table := table{
		name: tableName,
	}
	return &CustomeTransactionAccountTable{
		table: table,
		ID:    strings.Join([]string{transactionAccounts, idColumn}, "."),
		Name:  strings.Join([]string{transactionAccounts, nameColumn}, "."),
	}
}
