package validator

import (
	"regexp"

	"gopkg.in/go-playground/validator.v9"

	"go_worlder_system/errs"

	log "github.com/sirupsen/logrus"
)

var (
	validate = newValidate()
)

func newValidate() *validator.Validate {
	validate := validator.New()
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return regexp.MustCompile(`^[a-zA-Z\d]{8,16}$`).MatchString(password)
	})
	validate.RegisterValidation("transaction-partner-type", func(fl validator.FieldLevel) bool {
		ls := []string{
			"vendor",
			"payee",
		}
		for _, v := range ls {
			if v == fl.Field().String() {
				return true
			}
		}
		return false
	})
	return validate
}

// Struct ...
func Struct(data interface{}) error {
	err := validate.Struct(data)
	if err != nil {
		log.Println(err.Error())
		return errs.Invalidated.Wrap(err, err.Error())
	}
	return nil
}
