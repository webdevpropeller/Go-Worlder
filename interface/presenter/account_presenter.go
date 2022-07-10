package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// AccountPresenter ...
type AccountPresenter struct {
}

// NewAccountPresenter ...
func NewAccountPresenter() outputport.AccountOutputPort {
	return &AccountPresenter{}
}

// Index ...
func (presenter *AccountPresenter) Index(accountList []model.Account) []outputdata.Account {
	oAccountList := []outputdata.Account{}
	for _, account := range accountList {
		oAccount := presenter.convert(&account)
		oAccountList = append(oAccountList, *oAccount)
	}
	return oAccountList
}

// Show ...
func (presenter *AccountPresenter) Show(account *model.Account) *outputdata.Account {
	return presenter.convert(account)
}

// New ...
func (presenter *AccountPresenter) New(accountNameList []model.AccountName) []outputdata.AccountName {
	oAccountNameList := []outputdata.AccountName{}
	for _, accountName := range accountNameList {
		oAccountName := outputdata.AccountName{
			ID:   accountName.ID,
			Name: accountName.Name,
		}
		oAccountNameList = append(oAccountNameList, oAccountName)
	}
	return oAccountNameList
}

func (presenter *AccountPresenter) convert(account *model.Account) *outputdata.Account {
	oType := &outputdata.AccountType{
		ID:   account.Name.Type.ID,
		Name: account.Name.Type.Name,
	}
	oName := &outputdata.AccountName{
		ID:   account.Name.ID,
		Name: account.Name.Name,
	}
	return &outputdata.Account{
		ID:      account.ID,
		UserID:  account.User.ID,
		Type:    oType,
		Name:    oName,
		Balance: float64(*account.Balance),
	}
}
