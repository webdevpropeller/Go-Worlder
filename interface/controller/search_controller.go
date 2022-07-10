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

// searchParam ...
type searchParam struct {
	Type       string
	CategoryID string
	Keyword    string
}

// SearchController ...
type SearchController struct {
	inputport inputport.SearchInputPort
	param     *searchParam
}

// NewSearchController ...
func NewSearchController(sqlHandler database.SQLHandler) *SearchController {
	param := &searchParam{}
	initializeParam(param)
	return &SearchController{
		inputport: interactor.NewSearchInteractor(
			presenter.NewSearchPresenter(),
			database.NewProductDatabase(sqlHandler),
		),
		param: param,
	}
}

// SearchProduct ...
// @summary Display products search page
// @description Pass all product list to the front
// @tags Search
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param keyword query string true "keyword"
// @success 200 {array} outputdata.Product
// @failure 404 {string} string "Can't display product search page"
// @router /search/product [get]
func (controller *SearchController) SearchProduct(c Context) error {
	keyword := c.QueryParam(controller.param.Keyword)
	iSearch := &inputdata.Search{
		Keyword: keyword,
	}
	productList, err := controller.inputport.SearchProduct(iSearch)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, productList)
}
