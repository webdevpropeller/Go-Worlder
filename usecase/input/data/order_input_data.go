package inputdata

// Order ...
type Order struct {
	UserID       string
	PurchaseList []Purchase
}

// PurchaseRequest ...
type PurchaseRequest struct {
	PurchaseList []Purchase
}

// Purchase ...
type Purchase struct {
	ProductID string
	Quantity  int
}
