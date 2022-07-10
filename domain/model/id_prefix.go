package model

import (
	"reflect"
)

// IDPrefix ...
type IDPrefix struct {
	UserID        string
	CardID        string
	BrandID       string
	ProductID     string
	ProjectID     string
	InventoryID   string
	AccountID     string
	TransactionID string
	PayeeID       string
	OrderID       string
	MessageID     string
}

// NewIDPrefix ...
func NewIDPrefix() *IDPrefix {
	idPrefix := &IDPrefix{}
	rv := reflect.Indirect(reflect.ValueOf(idPrefix))
	for i := 0; i < rv.NumField(); i++ {
		sf := rv.Type().Field(i)
		rv.Field(i).SetString(sf.Name)
	}
	return idPrefix
}
