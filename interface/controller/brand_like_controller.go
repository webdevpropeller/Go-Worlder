package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"log"
	"net/http"
)

// BrandLikeController ...
type BrandLikeController struct {
	inputport inputport.BrandLikeInputPort
}

// NewBrandLikeController ...
func NewBrandLikeController(sqlHandler database.SQLHandler) *BrandLikeController {
	return &BrandLikeController{
		inputport: interactor.NewBrandLikeInteractor(
			presenter.NewBrandLikePresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewBrandDatabase(sqlHandler),
			database.NewBrandLikeDatabase(sqlHandler),
		),
	}
}

// Index ...
// @summary
// @description Get brand like list by user id
// @tags BrandLike
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand id"
// @success 200 {array} outputdata.Brand ""
// @failure 404 {string} string "Brand list is not found"
// @router /brands/{id}/like [get]
func (controller *BrandLikeController) Index(c Context) error {
	userID := c.Param(idParam)
	brandLikeList, err := controller.inputport.Index(userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	c.JSON(http.StatusOK, brandLikeList)
	return c.JSON(http.StatusOK, nil)
}

// Create ...
// @summary Register an brandLike
// @description Register an brandLike
// @tags BrandLike
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand id"
// @success 200
// @failure 409 {string} string "Can't register an brandLike"
// @router /brands/{id}/like [post]
func (controller *BrandLikeController) Create(c Context) error {
	userID := c.UserID()
	id := c.Param(idParam)
	err := controller.inputport.Create(id, userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Delete ...
// @summary Delete a brandLike
// @description Delete a brandLike got by the id
// @tags BrandLike
// @produce json
// @param Authorization header string true "jwt token"
// @param id path string true "Brand ID"
// @success 200
// @failure 409 {string} string "The user can't delete the brandLike"
// @router /brands/{id}/like [delete]
func (controller *BrandLikeController) Delete(c Context) error {
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
