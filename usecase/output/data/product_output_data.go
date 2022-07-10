package outputdata

// Product ...
type Product struct {
	ID          string
	User        *UserSimplified
	Brand       *BrandSimplified
	Category    *Category
	GenreID     uint
	Name        string
	Price       uint
	Image       string
	Description string
}

// ProductInventory ...
type ProductInventory struct {
	ProductID string
	Receiving int
	Shipping  int
	Disposal  int
	Stock     int
}

// ProductManagement ...
type ProductManagement struct {
	ProductID string
	Storage   string
	Memo      string
	Barcode   string
}
