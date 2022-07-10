package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"net/http"

	log "github.com/sirupsen/logrus"
)

// inventoryParam ...
type inventoryParam struct {
	Quantity string
	Storage  string
	Memo     string
	Barcode  string
}

// InventoryController ...
type InventoryController struct {
	inputport inputport.InventoryInputPort
	param     *inventoryParam
}

// NewInventoryController ...
func NewInventoryController(sqlHandler database.SQLHandler) *InventoryController {
	param := &inventoryParam{}
	initializeParam(param)
	return &InventoryController{
		inputport: interactor.NewInventoryInteractor(
			presenter.NewInventoryPresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewInventoryDatabase(sqlHandler),
			database.NewProductDatabase(sqlHandler),
		),
		param: param,
	}
}

// Index ...
// @summary Get inventory list
// @description Get inventory list by user id
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "Inventory list is not found"
// @router /inventory/list [get]
func (controller *InventoryController) Index(c Context) error {
	userID := c.UserID()
	oInventoryList, err := controller.inputport.IndexProduct(userID)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oInventoryList)
}

// Show ...
// @summary Display a product
// @description Display a product by the id
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Product ID"
// @success 200 {object} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The product is not found"
// @router /inventory/list/{id} [get]
func (controller *InventoryController) Show(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	oInventory, err := controller.inputport.ShowProduct(id, userID)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oInventory)
}

// New ...
func (controller *InventoryController) New(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Edit ...
// @summary Display a product
// @description Display a product by the id
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Product ID"
// @success 200 {object} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The product is not found"
// @router /inventory/list/{id}/edit [get]
func (controller *InventoryController) Edit(c Context) error {
	id := c.Param(idParam)
	userID := c.UserID()
	data, err := controller.inputport.EditProduct(id, userID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, data)
}

// Create ...
func (controller *InventoryController) Create(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Update ...
// @summary Display a product
// @description Display a product by the id
// @tags Inventory
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Product ID"
// @param storage formData string false "storage"
// @param memo formData string false "memo"
// @param barcode formData string false "barcode"
// @success 200 {object} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The product is not found"
// @router /inventory/list/{id} [patch]
func (controller *InventoryController) Update(c Context) error {
	productID := c.Param(idParam)
	userID := c.UserID()
	storage := c.FormValue(controller.param.Storage)
	memo := c.FormValue(controller.param.Memo)
	barcode := c.FormValue(controller.param.Barcode)
	iProductManagement := &inputdata.ProductManagement{
		ProductID: productID,
		Storage:   storage,
		Memo:      memo,
		Barcode:   barcode,
	}
	err := controller.inputport.UpdateProduct(iProductManagement, userID)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, iProductManagement)
}

// Delete ...
func (controller *InventoryController) Delete(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// NewReceiving ...
// @summary Display receiving page
// @description Display receiving page
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "Receiving list is not found"
// @router /inventory/receiving [get]
func (controller *InventoryController) NewReceiving(c Context) error {
	userID := c.UserID()
	productInventoryList, err := controller.inputport.New(userID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, productInventoryList)
}

// CreateReceiving ...
// @summary Create receiving list
// @description Create receiving list
// @tags Inventory
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @param updateList body inputdata.InventoryRequest true "Receiving request"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The user can't recieve"
// @router /inventory/receiving [post]
func (controller *InventoryController) CreateReceiving(c Context) error {
	userID := c.UserID()
	inventoryRequest := &inputdata.InventoryRequest{}
	err := c.Bind(inventoryRequest)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, errs.Cause(err).Error())
		return err
	}
	iInventory := &inputdata.Inventory{
		UserID:  userID,
		Type:    receivingType,
		Request: inventoryRequest,
	}
	err = controller.inputport.Create(iInventory)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// NewShipping ...
// @summary Get shipping list
// @description Get shipping list by user id
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "Shipping list is not found"
// @router /inventory/shipping [get]
func (controller *InventoryController) NewShipping(c Context) error {
	userID := c.UserID()
	productList, err := controller.inputport.New(userID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, productList)
}

// CreateShipping ...
// @summary Display receiving list of the user
// @description Display receiving list of the user
// @tags Inventory
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @param updateList body inputdata.InventoryRequest true "Receiving request"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The user can't recieve"
// @router /inventory/shipping [post]
func (controller *InventoryController) CreateShipping(c Context) error {
	userID := c.UserID()
	inventoryRequest := &inputdata.InventoryRequest{}
	err := c.Bind(inventoryRequest)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, errs.Cause(err).Error())
		return err
	}
	iInventory := &inputdata.Inventory{
		UserID:  userID,
		Type:    receivingType,
		Request: inventoryRequest,
	}
	err = controller.inputport.Create(iInventory)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// NewDisposal ...
// @summary Display receiving page
// @description Display receiving page
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "Disposal list is not found"
// @router /inventory/disposal [get]
func (controller *InventoryController) NewDisposal(c Context) error {
	userID := c.UserID()
	productList, err := controller.inputport.New(userID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, productList)
}

// CreateDisposal ...
// @summary Display receiving list of the user
// @description Display receiving list of the user
// @tags Inventory
// @accept json
// @produce json
// @param Authorization header string true "jwt token"
// @param updateList body inputdata.InventoryRequest true "Receiving request"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "The user can't recieve"
// @router /inventory/disposal [post]
func (controller *InventoryController) CreateDisposal(c Context) error {
	userID := c.UserID()
	inventoryRequest := &inputdata.InventoryRequest{}
	err := c.Bind(inventoryRequest)
	if err != nil {
		log.Error(err)
		c.String(http.StatusBadRequest, errs.Cause(err).Error())
		return err
	}
	iInventory := &inputdata.Inventory{
		UserID:  userID,
		Type:    receivingType,
		Request: inventoryRequest,
	}
	err = controller.inputport.Create(iInventory)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// NewStocktaking ...
// @summary Display receiving page
// @description Display receiving page
// @tags Inventory
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string "Stocktaking list is not found"
// @router /inventory/stocktaking [get]
func (controller *InventoryController) NewStocktaking(c Context) error {
	userID := c.UserID()
	productList, err := controller.inputport.New(userID)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, productList)
}

// CreateStocktaking ...
// @summary Display receiving list of the user
// @description Display receiving list of the user
// @tags Inventory
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param updateList body inputdata.StocktakingRequest true "Receiving request"
// @success 200 {array} outputdata.Inventory "Inventory"
// @failure 404 {string} string ""
// @router /inventory/stocktaking [post]
func (controller *InventoryController) CreateStocktaking(c Context) error {
	userID := c.UserID()
	iStocktakingRequest := &inputdata.StocktakingRequest{}
	c.Bind(iStocktakingRequest)
	iStocktaking := &inputdata.Stocktaking{
		UserID:  userID,
		Request: iStocktakingRequest,
	}
	err := controller.inputport.UpdateStocktaking(iStocktaking)
	if err != nil {
		log.Error(err)
		c.String(http.StatusNotFound, errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}
