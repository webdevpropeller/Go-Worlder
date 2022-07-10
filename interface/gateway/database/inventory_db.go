package database

// InventoryDB ...
type InventoryDB struct {
	db
	MInventoryTypesTable MInventoryTypesTable
	InventoryTable       InventoryTable
}

// MInventoryTypesTable ...
type MInventoryTypesTable struct {
	table
	ID   string
	Name string
}

// InventoryTable ...
type InventoryTable struct {
	table
	ID              string
	UserID          string
	ProductID       string
	InventoryTypeID string
	Quantity        string
	CreatedAt       string
	UpdatedAt       string
}

// NewInventoryDB ...
func NewInventoryDB() *InventoryDB {
	inventoryDB := &InventoryDB{}
	initialize(inventoryDB)
	return inventoryDB
}
