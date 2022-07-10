package model

import (
	"go_worlder_system/validator"
	"mime/multipart"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// Profile ...
type Profile struct {
	UserID     string `validate:"required"`
	Activity   string `validate:"required"`
	Industry   string `validate:"required"`
	Company    string `validate:"required"`
	Country    string
	Address1   string
	Address2   string
	ZipCode    string `validate:"number"`
	URL        string
	Phone      string `validate:"number"`
	AccountID  string `validate:"required"`
	LogoPath   string
	Logo       *multipart.FileHeader
	FirstName  string `validate:"required"`
	MiddleName string
	FamilyName string `validate:"required"`
	IconPath   string
	Icon       *multipart.FileHeader
}

// User is an entity for authentication
type User struct {
	ID             string `validate:"required"`
	Email          string `validate:"required,email"`
	PasswordDigest string `validate:"required"`
	Category       int    `validate:"required,min=1,max=2"`
	Profile        *Profile
}

// NewUser ...
func NewUser(email string, password string, category int) (*User, error) {
	id := generateID(idPrefix.UserID)
	passwordDigets, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user := &User{
		ID:             id,
		Email:          email,
		PasswordDigest: string(passwordDigets),
		Category:       category,
	}
	err = validator.Struct(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return user, nil
}

// IsValidPassword ...
func (user *User) IsValidPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordDigest), []byte(password+os.Getenv("WS_PASSWORD_SALT")))
	return err == nil
}

func (user *User) UpdatePassword(password string) error {
	passwordDigets, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Error(err)
		return err
	}
	user.PasswordDigest = string(passwordDigets)
	return nil
}

// CreateProfile ...
func (user *User) CreateProfile(
	activityID string,
	industryID string,
	company string,
	countryID string,
	address1 string,
	address2 string,
	zipCode string,
	url string,
	phone string,
	accountID string,
	logo *multipart.FileHeader,
	firstName string,
	middleName string,
	familyName string,
	icon *multipart.FileHeader,
) error {
	var logoPath string
	if logo != nil {
		logoPath = logo.Filename
	}
	var iconPath string
	if icon != nil {
		iconPath = logo.Filename
	}
	profile := &Profile{
		UserID:     user.ID,
		Activity:   activityID,
		Industry:   industryID,
		Company:    company,
		Country:    countryID,
		Address1:   address1,
		Address2:   address2,
		ZipCode:    zipCode,
		Phone:      phone,
		URL:        url,
		AccountID:  accountID,
		LogoPath:   logoPath,
		Logo:       logo,
		FirstName:  firstName,
		MiddleName: middleName,
		FamilyName: familyName,
		IconPath:   iconPath,
		Icon:       icon,
	}
	err := validator.Struct(profile)
	if err != nil {
		log.Error(err)
		return err
	}
	user.Profile = profile
	return nil
}

// Update ...
func (user *User) Update(
	activityID string,
	industryID string,
	company string,
	countryID string,
	address1 string,
	address2 string,
	zipCode string,
	url string,
	phone string,
	accountID string,
	logo *multipart.FileHeader,
	firstName string,
	middleName string,
	familyName string,
	icon *multipart.FileHeader,
) error {
	user.Profile.Activity = activityID
	user.Profile.Industry = industryID
	user.Profile.Company = company
	user.Profile.Country = countryID
	user.Profile.Address1 = address1
	user.Profile.Address2 = address2
	user.Profile.ZipCode = zipCode
	user.Profile.Phone = phone
	user.Profile.AccountID = accountID
	user.Profile.URL = url
	user.Profile.Logo = logo
	user.Profile.FirstName = firstName
	user.Profile.MiddleName = middleName
	user.Profile.FamilyName = familyName
	user.Profile.Icon = icon
	err := validator.Struct(user.Profile)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
