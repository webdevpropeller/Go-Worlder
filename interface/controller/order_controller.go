package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// orderParam ...
type orderParam struct {
	TotalAmount string
}

// OrderController ...
type OrderController struct {
	inputport inputport.OrderInputPort
	param     *orderParam
}

// NewOrderController ...
func NewOrderController(sqlHandler database.SQLHandler) *OrderController {
	param := &orderParam{}
	initializeParam(param)
	return &OrderController{
		inputport: interactor.NewOrderInteractor(
			database.NewUserDatabase(sqlHandler),
			database.NewOrderDatabase(sqlHandler),
			database.NewProductDatabase(sqlHandler),
		),
		param: param,
	}
}

// Create ...
// @summary Purchase products
// @description Purchase products
// @tags Order
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @param purchaseList body inputdata.PurchaseRequest true "Purchase request"
// @param total-amount formData number true "total amount"
// @success 200
// @failure 409 {string} string "The user can't purchase products"
// @router /api/orders [post]
func (controller *OrderController) Create(c Context) error {
	userID := c.UserID()
	iPurchaseRequest := inputdata.PurchaseRequest{}
	err := c.Bind(iPurchaseRequest)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	iOrder := &inputdata.Order{
		UserID:       userID,
		PurchaseList: iPurchaseRequest.PurchaseList,
	}
	err = controller.inputport.Create(iOrder)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iOrder)
}
