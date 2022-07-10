package presenter

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
)

// PayeePresenter ...
type PayeePresenter struct {
}

// NewPayeePresenter ...
func NewPayeePresenter() outputport.PayeeOutputPort {
	return &PayeePresenter{}
}

// Index ...
func (presenter *PayeePresenter) Index(payeeList []model.Payee) []outputdata.Payee {
	oPayeeList := []outputdata.Payee{}
	for _, payee := range payeeList {
		oPayee := presenter.convert(&payee)
		oPayeeList = append(oPayeeList, *oPayee)
	}
	return oPayeeList
}

// Show ...
func (presenter *PayeePresenter) Show(payee *model.Payee) *outputdata.Payee {
	return presenter.convert(payee)
}

// Edit ...
func (presenter *PayeePresenter) Edit(payee *model.Payee) *outputdata.Payee {
	return presenter.convert(payee)
}

func (presenter *PayeePresenter) convert(payee *model.Payee) *outputdata.Payee {
	return &outputdata.Payee{
		ID:   payee.ID,
		Name: payee.Name,
	}
}
