package database

// ProductDB ...
type ProductDB struct {
	db
	ProductsTable          ProductsTable
	ProductInventoryTable  ProductInventoryTable
	ProductManagementTable ProductManagementTable
	ProductsActiveTable    ProductsActiveTable
	ProductsFormalTable    ProductsFormalTable
	ProductsOfDraftTable   ProductsOfDraftTable
	ProductsDeletedTable   ProductsDeletedTable
}

// ProductsTable ...
type ProductsTable struct {
	table
	ID          string
	BrandID     string
	CategoryID  string
	GenreID     string
	Name        string
	Price       string
	Image       string
	Description string
	IsDraft     string
	CreatedAt   string
	UpdatedAt   string
}

// ProductInventoryTable ...
type ProductInventoryTable struct {
	table
	ProductID string
	Receiving string
	Shipping  string
	Disposal  string
	Stock     string
	CreatedAt string
	UpdatedAt string
}

// ProductManagementTable ...
type ProductManagementTable struct {
	table
	ProductID string
	Storage   string
	Memo      string
	Barcode   string
	CreatedAt string
	UpdatedAt string
}

// ProductsActiveTable ...
type ProductsActiveTable struct {
	table
	ProductID string
	CreatedAt string
}

// ProductsFormalTable ...
type ProductsFormalTable struct {
	table
	ProductID string
	CreatedAt string
}

// ProductsOfDraftTable ...
type ProductsOfDraftTable struct {
	table
	ProductID string
	CreatedAt string
}

// ProductsDeletedTable ...
type ProductsDeletedTable struct {
	table
	ProductID string
	CreatedAt string
}

// NewProductDB ...
func NewProductDB() *ProductDB {
	productDB := &ProductDB{}
	initialize(productDB)
	return productDB
}
