package controller

import (
	"go_worlder_system/str"
	"reflect"
)

type paramName struct {
	Activity   string
	Industry   string
	Company    string
	Country    string
	Address1   string
	Address2   string
	City       string
	State      string
	ZipCode    string
	URL        string
	Phone      string
	AccountID  string
	Logo       string
	Language   string
	FirstName  string
	MiddleName string
	FamilyName string
	Icon       string
	Card       string
	Jwt        string
}

func newParamName() *paramName {
	pn := &paramName{}
	rv := reflect.Indirect(reflect.ValueOf(pn))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		value := str.ToSnakeCase(sf.Name)
		rv.Field(i).SetString(value)
	}
	return pn
}
