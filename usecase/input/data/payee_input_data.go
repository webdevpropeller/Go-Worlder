package inputdata

// Payee ...
type Payee struct {
	UserID string
	Name   string
}

// NewPayee ...
type NewPayee struct {
	*Payee
}

// UpdatedPayee ...
type UpdatedPayee struct {
	ID string
	*Payee
}
