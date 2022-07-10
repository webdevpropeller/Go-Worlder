package outputdata

// AccountItem ...
type AccountItem struct {
	ID   string
	Name string
}

// TransactionPartner ...
type TransactionPartner struct {
	ID   string
	Name string
}

// TransactionAccount ...
type TransactionAccount struct {
	ID   string
	Name string
}

// Transaction ...
type Transaction struct {
	ID          string              `validate:"required"`
	User        *UserSimplified     `validate:"required"`
	AccountItem *AccountItem        `validate:"required"`
	Partner     *TransactionPartner `validate:"required"`
	Account     *TransactionAccount `validate:"required"`
	IsIncome    bool
	IsPaid      bool
	AccrualDate string  `validate:"required"`
	Amount      float64 `validate:"required"`
	Remarks     string
}
