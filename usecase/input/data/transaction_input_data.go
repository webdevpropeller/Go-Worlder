package inputdata

// Transaction ...
type Transaction struct {
	UserID        string `validate:"required"`
	IsIncome      bool
	IsPaid        bool
	AccountID     string  `validate:"required"`
	AccrualDate   string  `validate:"required"`
	AccountItemID string  `validate:"required"`
	PartnerID     string  `validate:"required"`
	Amount        float64 `validate:"required"`
	Remarks       string
}

// NewTransaction ...
type NewTransaction struct {
	*Transaction
}

// UpdatedTransaction ...
type UpdatedTransaction struct {
	ID string
	*Transaction
}
