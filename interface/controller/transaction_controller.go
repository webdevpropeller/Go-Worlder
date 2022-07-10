package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"net/http"
	"strconv"

	log "github.com/sirupsen/logrus"
)

// transactionParam ...
type transactionParam struct {
	IsIncome      string
	IsPaid        string
	AccountID     string
	AccrualDate   string
	AccountItemID string
	PartnerID     string
	Amount        string
	Remarks       string
	Type          string
}

// TransactionController ...
type TransactionController struct {
	inputport inputport.TransactionInputPort
	param     *transactionParam
}

// NewTransactionController ...
func NewTransactionController(sqlHandler database.SQLHandler) *TransactionController {
	param := &transactionParam{}
	initializeParam(param)
	return &TransactionController{
		inputport: interactor.NewTransactionInteractor(
			presenter.NewTransactionPresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewTransactionDatabase(sqlHandler),
			database.NewAccountDatabase(sqlHandler),
		),
		param: param,
	}
}

// Index ...
// @summary Display balance of payment management page
// @description Display balance of payment management page, getting transaction data from database
// @tags Transaction
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Transaction "transaction List"
// @failure 404 {string} string "Transactions is not found"
// @router /accounting/transactions [get]
func (controller *TransactionController) Index(c Context) error {
	userID := c.UserID()
	transactionList, err := controller.inputport.Index(userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, transactionList)
}

// Show ...
func (controller *TransactionController) Show(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// IndexAccount ...
// @summary Display balance of payment management page
// @description Display balance of payment management page, getting transaction data from database
// @tags Transaction
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Account ID"
// @success 200
// @failure 404 {string} string "Transactions is not found"
// @router /accounting/transactions/accounts/{id} [get]
func (controller *TransactionController) IndexAccount(c Context) error {
	userID := c.UserID()
	accountID := c.Param(idParam)
	transactionList, err := controller.inputport.IndexAccount(accountID, userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusNotFound, "Transaction list is not found")
		return err
	}
	return c.JSON(http.StatusOK, transactionList)
}

// New ...
// @summary
// @description
// @tags Transaction
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.AccountItem "accountItemList"
// @router /accounting/transactions/new [get]
func (controller *TransactionController) New(c Context) error {
	accountItemList, err := controller.inputport.New()
	if err != nil {
		log.Println(err)
		c.String(http.StatusNotFound, "Account item list is not found")
		return err
	}
	return c.JSON(http.StatusOK, accountItemList)
}

// Create ...
// @summary Register transaction
// @description Register transaction
// @tags Transaction
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param is-income formData boolean true "income or expenditure"
// @param is-paid formData boolean true "is-paid"
// @param account-id formData string true "account-id"
// @param accrual-date formData string true "accrual date"
// @param account-item-id formData string true "account item id, 1"
// @param partner-id formData string true "partner-id"
// @param amount formData number true "amount"
// @param remarks formData string false "remarks"
// @success 200
// @failure 409 {string} string "The user can't register transaction"
// @router /accounting/transactions [post]
func (controller *TransactionController) Create(c Context) error {
	userID := c.UserID()
	isIncome, err := strconv.ParseBool(c.FormValue(controller.param.IsIncome))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	isPaid, err := strconv.ParseBool(c.FormValue(controller.param.IsPaid))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	accountID := c.FormValue(controller.param.AccountID)
	accrualDate := c.FormValue(controller.param.AccrualDate)
	accountItemID := c.FormValue(controller.param.AccountItemID)
	partnerID := c.FormValue(controller.param.PartnerID)
	amount, err := strconv.ParseFloat(c.FormValue(controller.param.Amount), 64)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	remarks := c.FormValue(controller.param.Remarks)
	transaction := &inputdata.Transaction{
		UserID:        userID,
		IsIncome:      isIncome,
		IsPaid:        isPaid,
		AccountID:     accountID,
		AccrualDate:   accrualDate,
		AccountItemID: accountItemID,
		PartnerID:     partnerID,
		Amount:        amount,
		Remarks:       remarks,
	}
	iNewTransaction := &inputdata.NewTransaction{Transaction: transaction}
	err = controller.inputport.Create(iNewTransaction)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), "The user can't register transaction")
		return err
	}
	return c.JSON(http.StatusOK, transaction)
}

// Edit ...
// @summary Display edit page of a transaction
// @description Display edit page of a transaction got by the id
// @tags Transaction
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Transaction ID"
// @success 200 {object} outputdata.Transaction "transaction"
// @failure 404 {string} string "The user can't get the transaction"
// @router /accounting/transactions/{id}/edit [get]
func (controller *TransactionController) Edit(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	transaction, err := controller.inputport.Edit(id, userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusNotFound, "The user can't get the transaction")
		return err
	}
	return c.JSON(http.StatusOK, transaction)
}

// Update ...
// @summary Edit a transaction
// @description Edit a transaction
// @tags Transaction
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Transaction ID"
// @param is-income formData boolean true "income or expenditure"
// @param is-paid formData boolean true "is-paid"
// @param account-id formData string true "account-id"
// @param accrual-date formData string true "accrual date"
// @param account-item-id formData int true "account item id, 1"
// @param partner-id formData string true "partner-id"
// @param amount formData number true "amount"
// @param remarks formData string false "remarks"
// @success 200
// @failure 409 {string} string "The user can't edit a transaction"
// @router /accounting/transactions/{id} [patch]
func (controller *TransactionController) Update(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	isIncome, err := strconv.ParseBool(c.FormValue(controller.param.IsIncome))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	isPaid, err := strconv.ParseBool(c.FormValue(controller.param.IsPaid))
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	accountID := c.FormValue(controller.param.AccountID)
	accrualDate := c.FormValue(controller.param.AccrualDate)
	accountItemID := c.FormValue(controller.param.AccountItemID)
	partnerID := c.FormValue(controller.param.PartnerID)
	amount, err := strconv.ParseFloat(c.FormValue(controller.param.Amount), 64)
	if err != nil {
		log.Println(err)
		c.String(http.StatusBadRequest, validationError)
		return err
	}
	remarks := c.FormValue(controller.param.Remarks)
	iTransaction := &inputdata.Transaction{
		UserID:        userID,
		IsIncome:      isIncome,
		IsPaid:        isPaid,
		AccountID:     accountID,
		AccrualDate:   accrualDate,
		AccountItemID: accountItemID,
		PartnerID:     partnerID,
		Amount:        amount,
		Remarks:       remarks,
	}
	iUpdatedTransaction := &inputdata.UpdatedTransaction{
		ID:          id,
		Transaction: iTransaction,
	}
	err = controller.inputport.Update(iUpdatedTransaction)
	if err != nil {
		log.Println(err)
		c.String(http.StatusConflict, "The user can't edit a transaction")
		return err
	}
	return c.JSON(http.StatusOK, iUpdatedTransaction)
}

// Delete ...
// @summary Delete a transaction
// @description Delete a transaction got by the id
// @tags Transaction
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Transaction ID"
// @success 200
// @failure 409 {string} string "The user can't delete the transaction"
// @router /accounting/transactions/{id} [delete]
func (controller *TransactionController) Delete(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	err := controller.inputport.Delete(id, userID)
	if err != nil {
		log.Println(err)
		c.String(http.StatusConflict, "The user can't delete the transaction")
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
