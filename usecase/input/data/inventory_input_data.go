package inputdata

// ProductManagement ...
type ProductManagement struct {
	ProductID string
	Storage   string
	Memo      string
	Barcode   string
}

// Inventory ...
type Inventory struct {
	UserID  string
	Type    int
	Request *InventoryRequest
}

// InventoryRequest ...
type InventoryRequest struct {
	UpdateList []struct {
		ProductID string
		Quantity  int
	}
}

// Stocktaking ...
type Stocktaking struct {
	UserID  string
	Request *StocktakingRequest
}

// StocktakingRequest ...
type StocktakingRequest struct {
	UpdateList []struct {
		ProductID string
		Stock     int
	}
}
