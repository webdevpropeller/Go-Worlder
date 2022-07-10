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

// payeeParam ...
type payeeParam struct {
	Name string
}

// PayeeController ...
type PayeeController struct {
	inputport inputport.PayeeInputPort
	param     *payeeParam
}

// NewPayeeController ...
func NewPayeeController(sqlHandler database.SQLHandler) *PayeeController {
	param := &payeeParam{}
	initializeParam(param)
	return &PayeeController{
		inputport: interactor.NewPayeeInteractor(
			presenter.NewPayeePresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewPayeeDatabase(sqlHandler),
		),
		param: param,
	}
}

// Index ...
// @summary Display payee list
// @description Display payee list
// @tags Payee
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Payee "Payee list"
// @failure 404 {string} string "The user can't get payee list"
// @router /accounting/transactions/payees [get]
func (controller *PayeeController) Index(c Context) error {
	userID := c.UserID()
	payeeList, err := controller.inputport.Index(userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, payeeList)
}

// Create ...
// @summary Register a payee
// @description Register a payee
// @tags Payee
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param name formData string false "remarks"
// @success 200
// @failure 409 {string} string "The user can't create a payee"
// @router /accounting/transactions/payees [post]
func (controller *PayeeController) Create(c Context) error {
	userID := c.UserID()
	name := c.FormValue(controller.param.Name)
	iPayee := &inputdata.Payee{
		UserID: userID,
		Name:   name,
	}
	iNewPayee := &inputdata.NewPayee{Payee: iPayee}
	err := controller.inputport.Create(iNewPayee)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iPayee)
}

// Update ...
// @summary Edit payee
// @description Edit a payee got by the payee id
// @tags Payee
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "payee id"
// @param name formData string true "payee name"
// @success 200 {object} outputdata.Payee "payee"
// @failure 409 {string} string "The user can't update the payee"
// @router /accounting/transactions/payees/{id} [patch]
func (controller *PayeeController) Update(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	name := c.FormValue(controller.param.Name)
	iPayee := &inputdata.Payee{
		UserID: userID,
		Name:   name,
	}
	iUpdatedPayee := &inputdata.UpdatedPayee{
		ID:    id,
		Payee: iPayee,
	}
	err := controller.inputport.Update(iUpdatedPayee)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iPayee)
}

// Delete ...
// @summary Delete a payee
// @description Delete a payee
// @tags Payee
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Payee ID"
// @success 200
// @failure 409 {string} string "The user can't create a payee"
// @router /accounting/transactions/payees/{id} [delete]
func (controller *PayeeController) Delete(c Context) error {
	id := c.UserID()
	userID := c.UserID()
	err := controller.inputport.Delete(id, userID)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
