package controller

import (
	"net/http"
	"reflect"

	"go_worlder_system/errs"
	"go_worlder_system/str"
)

const (
	receivingType int = 1
	shippingType  int = 2
	disposalType  int = 3
)

var (
	pn = newParamName()
	// context values
	idParam = "id"
	// error messages
	validationError = "Validation error"
)

func statusCode(err error) int {
	switch errs.GetType(err) {
	case errs.Unknown:
		return http.StatusBadRequest
	case errs.Invalidated:
		return http.StatusBadRequest
	case errs.Unauthorized:
		return http.StatusUnauthorized
	case errs.Forbidden:
		return http.StatusForbidden
	case errs.NotFound:
		return http.StatusNotFound
	case errs.Conflict:
		return http.StatusConflict
	case errs.Failed:
		return http.StatusInternalServerError
	default:
		return http.StatusBadRequest
	}
}

func initializeParam(param interface{}) {
	rv := reflect.Indirect(reflect.ValueOf(param))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		value := str.ToKebabCase(sf.Name)
		rv.Field(i).SetString(value)
	}
	return
}
