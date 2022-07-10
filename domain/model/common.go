package model

import (
	"go_worlder_system/str"
	"regexp"
	"strings"

	"gopkg.in/go-playground/validator.v9"
)

const (
	idLength    int = 32
	tokenLength int = 32
	// inventory type id
	receivingType   int = 1
	shippingType    int = 2
	faultyType      int = 3
	stocktakingType int = 4
)

var (
	validate   = newValidate()
	rs1Letters = []rune("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	// id prefix
	idPrefix = NewIDPrefix()
)

func newValidate() (validate *validator.Validate) {
	validate = validator.New()
	validate.RegisterValidation("password", func(fl validator.FieldLevel) bool {
		password := fl.Field().String()
		return regexp.MustCompile(`^[a-zA-Z\d]{8,16}$`).MatchString(password)
	})
	return
}

func generateID(idPrefix string) string {
	prefix := strings.Join([]string{idPrefix, "_"}, "")
	id := prefix + str.Random(idLength-len(prefix))
	return id
}
