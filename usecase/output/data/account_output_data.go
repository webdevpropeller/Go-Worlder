package outputdata

// AccountType ...
type AccountType struct {
	ID   int
	Name string
}

// AccountName ...
type AccountName struct {
	ID   string
	Name string
}

// AccountBalance ...
type AccountBalance struct {
	RegisteredBalance   float64
	SynchronizedBalance float64
	UnregisteredCosts   float64
}

// Account ...
type Account struct {
	ID      string
	UserID  string
	Type    *AccountType
	Name    *AccountName
	Balance float64
}

// AccountSimplified ...
type AccountSimplified struct {
	ID   string
	Name string
}
