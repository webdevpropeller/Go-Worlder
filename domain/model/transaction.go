package model

import (
	"go_worlder_system/errs"
	"go_worlder_system/validator"

	log "github.com/sirupsen/logrus"
)

// AccountItem ...
type AccountItem struct {
	ID   string `validate:"required"`
	Name string `validate:"required"`
}

// TransactionPartner ...
type TransactionPartner struct {
	ID   string
	Name string
	Type string `validate:"required,transaction-partner-type"`
}

// TransactionAccount ...
type TransactionAccount struct {
	ID   string
	Name string
}

// Transaction is an entity
type Transaction struct {
	ID          string              `validate:"required"`
	User        *User               `validate:"required"`
	AccountItem *AccountItem        `validate:"required"`
	Partner     *TransactionPartner `validate:"required"`
	Account     *TransactionAccount
	IsIncome    bool
	IsPaid      bool
	AccrualDate string  `validate:"required"`
	Amount      float64 `validate:"required"`
	Remarks     string
}

// TransactionAccountOption ...
type TransactionAccountOption func(*Transaction)

// TransactionByAccount ...
func TransactionByAccount(account *Account) TransactionAccountOption {
	return func(transaction *Transaction) {
		transaction.Account = &TransactionAccount{
			ID:   account.ID,
			Name: account.Name.Name,
		}
	}
}

// TransactionByCash ...
func TransactionByCash(cash *Cash) TransactionAccountOption {
	return func(transaction *Transaction) {
		transaction.Account = &TransactionAccount{
			ID:   "cash",
			Name: "cash",
		}
	}
}

// NewTransaction ...
func NewTransaction(
	user *User,
	accountItem *AccountItem,
	partner *TransactionPartner,
	accountOption TransactionAccountOption,
	isIncome bool,
	isPaid bool,
	accrualDate string,
	amount float64,
	remarks string,
) (*Transaction, error) {
	id := generateID(idPrefix.TransactionID)
	transaction := &Transaction{
		ID:          id,
		User:        user,
		AccountItem: accountItem,
		Partner:     partner,
		IsIncome:    isIncome,
		IsPaid:      isPaid,
		AccrualDate: accrualDate,
		Amount:      amount,
		Remarks:     remarks,
	}
	accountOption(transaction)
	err := validator.Struct(transaction)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return transaction, nil
}

// IsOwner ...
func (transaction *Transaction) IsOwner(userID string) bool {
	return transaction.User.ID == userID
}

// Update ...
func (transaction *Transaction) Update(
	userID string,
	accountItem *AccountItem,
	partner *TransactionPartner,
	accountOption TransactionAccountOption,
	isIncome bool,
	isPaid bool,
	accrualDate string,
	amount float64,
	remarks string,
) error {
	if transaction.User.ID != userID {
		return errs.Forbidden.New("The user is not owner")
	}
	transaction.AccountItem = accountItem
	transaction.Partner = partner
	transaction.IsIncome = isIncome
	transaction.IsPaid = isPaid
	transaction.AccrualDate = accrualDate
	transaction.Amount = amount
	transaction.Remarks = remarks
	accountOption(transaction)
	err := validator.Struct(transaction)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
