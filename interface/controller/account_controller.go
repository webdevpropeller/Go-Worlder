package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"log"
	"net/http"
)

// accountParam ...
type accountParam struct {
	AccountTypeID string
	AccountNameID string
}

// AccountController ...
type AccountController struct {
	inputport inputport.AccountInputPort
	param     *accountParam
}

// NewAccountController ...
func NewAccountController(sqlHandler database.SQLHandler) *AccountController {
	param := &accountParam{}
	initializeParam(param)
	return &AccountController{
		inputport: interactor.NewAccountInteractor(
			presenter.NewAccountPresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewAccountDatabase(sqlHandler),
			database.NewAccountTypeDatabase(sqlHandler),
			database.NewAccountNameDatabase(sqlHandler),
		),
		param: param,
	}
}

// Index ...
// @summary
// @description Get account list by user id
// @tags Account
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Account ""
// @failure 404 {string} string "Account list is not found"
// @router /accounting/accounts [get]
func (controller *AccountController) Index(c Context) error {
	userID := c.UserID()
	oAccountList, err := controller.inputport.Index(userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oAccountList)
}

// Show ...
// @summary
// @description Get an account by id
// @tags Account
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Account ID"
// @success 200 {object} outputdata.Account "account"
// @failure 404 {string} string "The user can't get the account"
// @router /accounting/accounts/{id} [get]
func (controller *AccountController) Show(c Context) error {
	userID := c.UserID()
	id := c.Param(idParam)
	account, err := controller.inputport.Show(id, userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, account)
}

// New ...
// @summary
// @description
// @tags Account
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.AccountName "accountNameList"
// @router /accounting/accounts/new [get]
func (controller *AccountController) New(c Context) error {
	accountNameList, err := controller.inputport.New()
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, accountNameList)
}

// Edit ...
func (controller *AccountController) Edit(c Context) error {
	return nil
}

// Create ...
// @summary Register an account
// @description Register an account
// @tags Account
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param account-name-id formData string true "account name id"
// @success 200
// @failure 409 {string} string "The user can't register an account"
// @router /accounting/accounts [post]
func (controller *AccountController) Create(c Context) error {
	userID := c.UserID()
	accountNameID := c.FormValue(controller.param.AccountNameID)
	iAccount := &inputdata.Account{
		UserID:        userID,
		AccountNameID: accountNameID,
	}
	err := controller.inputport.Create(iAccount)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iAccount)
}

// Delete ...
// @summary Delete a account
// @description Delete a account got by the id
// @tags Account
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Account ID"
// @success 200
// @failure 409 {string} string "The user can't delete the account"
// @router /accounting/accounts/{id} [delete]
func (controller *AccountController) Delete(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	err := controller.inputport.Delete(id, userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
