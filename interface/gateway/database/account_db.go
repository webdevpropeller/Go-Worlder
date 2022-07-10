package database

// MAccountTypesTable ...
type MAccountTypesTable struct {
	table
	ID   string
	Name string
}

// MAccountNamesTable ...
type MAccountNamesTable struct {
	table
	ID            string
	AccountTypeID string
	Name          string
}

// AccountsTable ...
type AccountsTable struct {
	table
	ID            string
	UserID        string
	AccountNameID string
	CreatedAt     string
	UpdatedAt     string
}

// AccountsActiveTable ...
type AccountsActiveTable struct {
	table
	AccountID string
	CreatedAt string
}

// AccountsDeletedTable ...
type AccountsDeletedTable struct {
	table
	AccountID string
	CreatedAt string
}

// AccountDB ...
type AccountDB struct {
	db
	MAccountTypesTable   MAccountTypesTable
	MAccountNamesTable   MAccountNamesTable
	AccountsTable        AccountsTable
	AccountsActiveTable  AccountsActiveTable
	AccountsDeletedTable AccountsDeletedTable
}

// NewAccountDB ...
func NewAccountDB() *AccountDB {
	accountDB := &AccountDB{}
	initialize(accountDB)
	return accountDB
}
