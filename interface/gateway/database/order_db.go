package database

// OrderDB ...
type OrderDB struct {
	db
	OrdersTable            OrdersTable
	ProductsPurchasedTable ProductsPurchasedTable
}

// OrdersTable ...
type OrdersTable struct {
	table
	ID        string
	UserID    string
	ProductID string
	Quantity  string
	CreatedAt string
	UpdatedAt string
}

// ProductsPurchasedTable ...
type ProductsPurchasedTable struct {
	table
	OrderID   string
	ProductID string
	Quantity  string
}

// NewOrderDB ...
func NewOrderDB() *OrderDB {
	orderDB := &OrderDB{}
	initialize(orderDB)
	return orderDB
}
