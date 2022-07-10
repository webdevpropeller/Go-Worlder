package controller

import (
	"encoding/json"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"go_worlder_system/validator"
	"net/http"
	"strconv"
	"strings"
	"time"

	log "github.com/sirupsen/logrus"
)

type CardData struct {
	Company      string `json:"Company"`
	Owner        string `json:"Owner"`
	Number1      string `json:"Number1"`
	Number2      string `json:"Number2"`
	Number3      string `json:"Number3"`
	Number4      string `json:"Number4"`
	Month        string `json:"Month"`
	Year         string `json:"Year"`
	SecurityCode string `json:"SecurityCode"`
}

// ProfileController ...
type ProfileController struct {
	inputport inputport.ProfileInputPort
}

// NewProfileController ...
func NewProfileController(sqlHandler database.SQLHandler) *ProfileController {
	return &ProfileController{
		inputport: interactor.NewProfileInteractor(
			presenter.NewProfilePresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewOptionDatabase(sqlHandler),
			database.NewBrandDatabase(sqlHandler),
			database.NewProductDatabase(sqlHandler),
			database.NewProjectDatabase(sqlHandler),
		),
	}
}

// Index ...
func (controller *ProfileController) Index(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// Show ...
// @summary Display a profile
// @description Get the user and his profile by cookie value and display profile page
// @tags Profile
// @produce json
// @success 200 {object} outputdata.PublicUser "Profile"
// @failure 404 {string} string "Profile is not found"
// @router /profile/:id [get]
func (controller *ProfileController) Show(c Context) error {
	userID := c.Param(idParam)
	oUser, err := controller.inputport.Show(userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oUser)
}

// New ...
// @summary Display user infomation register page
// @description Display user infomation register page
// @tags Profile
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {object} outputdata.ProfileSelectItem "User infomation select item"
// @router /profile/new [get]
func (controller *ProfileController) New(c Context) error {
	oProfileSelectItem, err := controller.inputport.New()
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oProfileSelectItem)
}

// Edit ...
// @summary Display a profile edit page
// @description Get the user and his profile by cookie value and display profile edit page
// @tags Profile
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {object} outputdata.Profile "User infomation"
// @failure 404 {string} string "User infomation is not found"
// @router /profile/edit [get]
func (controller *ProfileController) Edit(c Context) error {
	userID := c.UserID()
	profile, err := controller.inputport.Edit(userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, profile)
}

// Create ...
// @summary Create user infomation
// @description Get the user by cookie value and update user infomation of the user
// @tags Profile
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param activity formData string true "id:1"
// @param industry formData string true "id:1"
// @param company formData string true "company"
// @param country formData string true "id:1"
// @param address1 formData string false "address1"
// @param address2 formData string false "address2"
// @param zip_code formData integer false "zip code"
// @param url formData string false "url"
// @param phone formData string false "phone"
// @param account_id formData string true "account id"
// @param logo formData file false "icon image"
// @param first_name formData string false "first name"
// @param middle_name formData string false "middle name"
// @param family_name formData string false "family name"
// @param icon formData file false "icon"
// @param card formData string true "json:Company,Owner,Number1,Number2,Number3,Number4,Month,Year,SecurityCode"
// @success 200 {object} outputdata.Profile "User infomation"
// @failure 400 {string} string "Validation error"
// @failure	409 {string} string "Can't register user infomation"
// @router /profile [post]
func (controller *ProfileController) Create(c Context) error {
	userID := c.UserID()
	// card
	cardData := new(CardData)
	card := c.FormValue(pn.Card)
	if err := json.Unmarshal([]byte(card), cardData); err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	cardMonth, _ := strconv.Atoi(cardData.Month)
	cardYear, _ := strconv.Atoi(cardData.Year)
	expiry := time.Date(2000+cardYear, time.Month(cardMonth+1), 0, 0, 0, 0, 0, time.Local)
	iCard := &inputdata.Card{
		Company:      cardData.Company,
		Name:         cardData.Owner,
		Number:       strings.Join([]string{cardData.Number1, cardData.Number2, cardData.Number3, cardData.Number4}, ""),
		Expiry:       &expiry,
		SecurityCode: cardData.SecurityCode,
	}
	err := validator.Struct(iCard)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	activity := c.FormValue(pn.Activity)
	industry := c.FormValue(pn.Industry)
	company := c.FormValue(pn.Company)
	country := c.FormValue(pn.Country)
	address1 := c.FormValue(pn.Address1)
	address2 := c.FormValue(pn.Address2)
	zipCode := c.FormValue(pn.ZipCode)
	url := c.FormValue(pn.URL)
	phone := c.FormValue(pn.Phone)
	accountID := c.FormValue(pn.AccountID)
	logo, _ := c.FormFile(pn.Logo)
	firstName := c.FormValue(pn.FirstName)
	middleName := c.FormValue(pn.MiddleName)
	familyName := c.FormValue(pn.FamilyName)
	icon, _ := c.FormFile(pn.Icon)
	iProfile := &inputdata.Profile{
		UserID:     userID,
		ActivityID: activity,
		IndustryID: industry,
		Company:    company,
		CountryID:  country,
		Address1:   address1,
		Address2:   address2,
		ZipCode:    zipCode,
		URL:        url,
		Phone:      phone,
		AccountID:  accountID,
		Logo:       logo,
		FirstName:  firstName,
		MiddleName: middleName,
		FamilyName: familyName,
		Icon:       icon,
		Card:       *iCard,
	}
	err = validator.Struct(iProfile)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	oUser, err := controller.inputport.Create(iProfile)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oUser)
}

// Update ...
// @summary Update user infomation
// @description Update user infomation
// @tags Profile
// @accept mpfd
// @produce json
// @param Authorization header string true "jwt token"
// @param activity formData string true "id:1"
// @param industry formData string true "id:1"
// @param company formData string true "company"
// @param country formData string true "id:1"
// @param address1 formData string false "address1"
// @param address2 formData string false "address2"
// @param zip_code formData integer false "zip code"
// @param url formData string false "url"
// @param phone formData string false "phone"
// @param account_id formData string true "account id"
// @param logo formData file false "icon image"
// @param first_name formData string false "first name"
// @param middle_name formData string false "middle name"
// @param family_name formData string false "family name"
// @param icon formData file false "icon"
// @param card formData string true "json:Company,Owner,Number1,Number2,Number3,Number4,Month,Year,SecurityCode"
// @success 200
// @failure 400 {string} string "Validation error"
// @failure 404 {string} string "User infomation is not found"
// @router /profile [patch]
func (controller *ProfileController) Update(c Context) error {
	userID := c.UserID()
	// card
	cardData := new(CardData)
	card := c.FormValue(pn.Card)
	if err := json.Unmarshal([]byte(card), cardData); err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	cardMonth, _ := strconv.Atoi(cardData.Month)
	cardYear, _ := strconv.Atoi(cardData.Year)
	expiry := time.Date(2000+cardYear, time.Month(cardMonth+1), 0, 0, 0, 0, 0, time.Local)
	iCard := &inputdata.Card{
		Company:      cardData.Company,
		Name:         cardData.Owner,
		Number:       strings.Join([]string{cardData.Number1, cardData.Number2, cardData.Number3, cardData.Number4}, ""),
		Expiry:       &expiry,
		SecurityCode: cardData.SecurityCode,
	}
	activity := c.FormValue(pn.Activity)
	industry := c.FormValue(pn.Industry)
	company := c.FormValue(pn.Company)
	country := c.FormValue(pn.Country)
	address1 := c.FormValue(pn.Address1)
	address2 := c.FormValue(pn.Address2)
	zipCode := c.FormValue(pn.ZipCode)
	url := c.FormValue(pn.URL)
	phone := c.FormValue(pn.Phone)
	accountID := c.FormValue(pn.AccountID)
	logo, _ := c.FormFile(pn.Logo)
	firstName := c.FormValue(pn.FirstName)
	middleName := c.FormValue(pn.MiddleName)
	familyName := c.FormValue(pn.FamilyName)
	icon, _ := c.FormFile(pn.Icon)
	iProfile := &inputdata.Profile{
		UserID:     userID,
		ActivityID: activity,
		IndustryID: industry,
		Company:    company,
		CountryID:  country,
		Address1:   address1,
		Address2:   address2,
		ZipCode:    zipCode,
		URL:        url,
		Phone:      phone,
		AccountID:  accountID,
		Logo:       logo,
		FirstName:  firstName,
		MiddleName: middleName,
		FamilyName: familyName,
		Icon:       icon,
		Card:       *iCard,
	}
	err := validator.Struct(iProfile)
	if err != nil {
		log.Error(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	err = controller.inputport.Update(iProfile)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Delete ...
func (controller *ProfileController) Delete(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// IndexBrandLike ...
// @summary Display brand like
// @description Display brand like got by the user id
// @tags Profile
// @produce json
// @param Authorization header string true "jwt token"
// @success 200 {array} outputdata.Brand "Brand like"
// @failure 404 {string} string "Brand like is not found"
// @router /profile/brandlike [get]
func (controller *ProfileController) IndexBrandLike(c Context) error {
	userID := c.UserID()
	brandLikeList, err := controller.inputport.IndexBrandLike(userID)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, brandLikeList)
}
