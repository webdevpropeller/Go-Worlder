package inputdata

import (
	"mime/multipart"
	"time"
)

// Profile ...
type Profile struct {
	UserID     string `validate:"required"`
	ActivityID string `validate:"required"`
	IndustryID string `validate:"required"`
	Company    string `validate:"required"`
	CountryID  string
	Address1   string
	Address2   string
	ZipCode    string `validate:"number"`
	URL        string
	Phone      string `validate:"number"`
	AccountID  string `validate:"required"`
	Logo       *multipart.FileHeader
	FirstName  string `validate:"required"`
	MiddleName string
	FamilyName string `validate:"required"`
	Icon       *multipart.FileHeader
	Card       Card
}

type Card struct {
	Company      string     `validate:"required"`
	Name         string     `validate:"required"`
	Number       string     `validate:"required,len=16"`
	Expiry       *time.Time `validate:"required"`
	SecurityCode string     `validate:"required,min=3,max=4"`
}

type Address struct {
	Country  string
	Address1 string
	Address2 string
	City     string
	State    string
	ZipCode  string
}
