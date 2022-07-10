package database

// MAccountItemsTable ...
type MAccountItemsTable struct {
	table
	ID   string
	Name string
}

// TransactionPartnersTable ...
type TransactionPartnersTable struct {
	table
	ID string
}

// TransactionAccountsTable ...
type TransactionAccountsTable struct {
	table
	ID string
}

// PayeesTable ...
type PayeesTable struct {
	table
	ID        string
	UserID    string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// TransactionsTable ...
type TransactionsTable struct {
	table
	ID                   string
	UserID               string
	AccountItemID        string
	TransactionPartnerID string
	TransactionAccountID string
	IsIncome             string
	IsPaid               string
	AccrualDate          string
	Amount               string
	Remarks              string
	CreatedAt            string
	UpdatedAt            string
}

// TransactionDB ...
type TransactionDB struct {
	db
	MAccountItemsTable       MAccountItemsTable
	TransactionPartnersTable TransactionPartnersTable
	PayeesTable              PayeesTable
	TransactionAccountsTable TransactionAccountsTable
	TransactionsTable        TransactionsTable
}

// NewTransactionDB ...
func NewTransactionDB() *TransactionDB {
	transactionDB := &TransactionDB{}
	initialize(transactionDB)
	return transactionDB
}
