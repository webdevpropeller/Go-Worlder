package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// TransactionPresenter ...
type TransactionPresenter struct {
}

// NewTransactionPresenter ...
func NewTransactionPresenter() outputport.TransactionOutputPort {
	return &TransactionPresenter{}
}

// Index ...
func (presenter TransactionPresenter) Index(transactionList []model.Transaction) []outputdata.Transaction {
	oTransactionList := []outputdata.Transaction{}
	for _, transaction := range transactionList {
		oTransaction := presenter.convert(&transaction)
		oTransactionList = append(oTransactionList, *oTransaction)
	}
	return oTransactionList
}

// Show ...
func (presenter *TransactionPresenter) Show(transaction *model.Transaction) *outputdata.Transaction {
	return presenter.convert(transaction)
}

// Edit ...
func (presenter *TransactionPresenter) Edit(transaction *model.Transaction) *outputdata.Transaction {
	return presenter.convert(transaction)
}

func (presenter *TransactionPresenter) convert(transaction *model.Transaction) *outputdata.Transaction {
	user := &outputdata.UserSimplified{
		ID:   transaction.User.ID,
		Name: transaction.User.Profile.Company,
	}
	accountItem := &outputdata.AccountItem{
		ID:   transaction.AccountItem.ID,
		Name: transaction.AccountItem.Name,
	}
	account := &outputdata.TransactionAccount{
		ID:   transaction.Account.ID,
		Name: transaction.Account.Name,
	}
	partner := &outputdata.TransactionPartner{
		ID:   transaction.Partner.ID,
		Name: transaction.Partner.Name,
	}
	return &outputdata.Transaction{
		ID:          transaction.ID,
		User:        user,
		AccountItem: accountItem,
		Partner:     partner,
		Account:     account,
		IsIncome:    transaction.IsIncome,
		IsPaid:      transaction.IsPaid,
		AccrualDate: transaction.AccrualDate,
		Amount:      transaction.Amount,
		Remarks:     transaction.Remarks,
	}
}
