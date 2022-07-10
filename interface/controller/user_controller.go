package controller

import (
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/communication"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/presenter"
	inputdata "go_worlder_system/usecase/input/data"
	inputport "go_worlder_system/usecase/input/port"
	"go_worlder_system/usecase/interactor"
	"net/http"
	"os"

	log "github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

// userParam ...
type userParam struct {
	Email          string
	Password       string
	UserCategoryID string
	Token          string
	KeepLogin      string
}

// UserController ...
type UserController struct {
	inputport inputport.UserInputPort
	param     *userParam
}

// NewUserController passes sqlcontroller to usecase
func NewUserController(sqlHandler database.SQLHandler) *UserController {
	param := &userParam{}
	initializeParam(param)
	return &UserController{
		inputport: interactor.NewUserInteractor(
			presenter.NewUserPresenter(),
			database.NewUserDatabase(sqlHandler),
			database.NewUserTokenDatabase(sqlHandler),
		),
		param: param,
	}
}

// Home ...
// @summary Display home page
// @description Display home page if the user is not logged in
// @produce json
// @success 200
// @router / [get]
func (ctrl *UserController) Home(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// New ...
// @summary Display signup page
// @description Display sign up page if the user is not logged in
// @tags UserAuth
// @produce json
// @success 200
// @router /signup [get]
func (ctrl *UserController) New(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// SignUp receives form from signup page and send authentication url to a user
// @summary Provisionally sign up
// @description Redirect to signup page if email overlaps with an existing email
// @tags UserAuth
// @accept mpfd
// @produce json
// @param email formData string true "email"
// @param password formData string true "password"
// @success 200
// @failure 400 {string} string "Validation error"
// @failure 409 {string} string "The user can't sign up"
// @router /signup [post]
func (ctrl *UserController) SignUp(c Context) error {
	// Get formData
	email := c.FormValue(ctrl.param.Email)
	password := c.FormValue(ctrl.param.Password)
	brandOwner := 1
	// Create a user entity
	iUser := &inputdata.User{
		Email:    email,
		Password: password,
		Category: brandOwner,
	}
	iSignUp := &inputdata.SignUp{User: iUser}
	oSignUp, err := ctrl.inputport.SignUp(iSignUp)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	// Send authentication email to a user
	err = communication.SendMail(oSignUp.Email, "Sign up", oSignUp.Message)
	if err != nil {
		c.String(statusCode(err), oSignUp.Message)
		return err
	}
	return c.JSON(http.StatusOK, oSignUp.Message)
}

// Activate acquire a provisionally sign-up user from authentication URL
// and completely sign up the user
// @summary Completely sign up
// @description Make the user active if the authentication token of the URL sent to the user's mail exists in the DB
// @tags UserAuth
// @accept mpfd
// @produce json
// @param token query string true "token"
// @success 200
// @failure 404 {string} string "generate error message 'Token is invalid'"
// @router /activate [post]
func (ctrl *UserController) Activate(c Context) error {
	token := c.QueryParam(ctrl.param.Token)
	err := ctrl.inputport.Activate(token)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// NewSignin ...
// @summary Display signin page
// @description Display my page if the user is logged in
// @tags UserAuth
// @produce json
// @success 200
// @router /signin [get]
func (ctrl *UserController) NewSignin(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// SignIn ...
// @summary Sign in
// @description Generate cookie for login discrimination if email and password match DB. Otherwise redirect to sign in page.
// @tags UserAuth
// @accept mpfd
// @produce json
// @param email formData string true "email"
// @param password formData string true "password"
// @success 200 {object} outputdata.SignIn "IsSigninClear and JwtToken"
// @failure 409 {string} string "Email or Password is incorrect"
// @router /signin [post]
func (ctrl *UserController) SignIn(c Context) error {
	email := c.FormValue(ctrl.param.Email)
	password := c.FormValue(ctrl.param.Password)
	iSignIn := &inputdata.SignIn{
		Email:    email,
		Password: password,
	}
	oSignIn, err := ctrl.inputport.SignIn(iSignIn)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, oSignIn)
}

// ForgotPassword ...
// @summary Send Email with url for authentication
// @description Send URL with authentication token to display password reset page to the email if the email is active. Otherwise redirect to forgot password page.
// @tags UserAuth
// @accept mpfd
// @produce json
// @param email formData string true "email as userID"
// @success 200
// @failure 400 {string} string "The user can't send password reset email"
// @router /password/forgot [post]
func (ctrl *UserController) ForgotPassword(c Context) error {
	email := c.FormValue(ctrl.param.Email)
	iForgotPassword := &inputdata.ForgotPassword{
		Email: email,
	}
	oForgotPassword, err := ctrl.inputport.ForgotPassword(iForgotPassword)
	if err != nil {
		log.Println(err)
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	err = communication.SendMail(oForgotPassword.Email, "Password Reset", oForgotPassword.Message)
	if err != nil {
		c.String(statusCode(err), oForgotPassword.Message)
		return err
	}
	return c.JSON(http.StatusOK, oForgotPassword)
}

// ResetPassword ...
// @summary Reset a password
// @description Get the user by token and reset password of the user.
// @tags UserAuth
// @accept mpfd
// @produce json
// @param token formData string true "token"
// @param password formData string true "password"
// @success 200
// @failure	409 {string} string "The user can't reset the password"
// @router /password/reset [post]
func (ctrl *UserController) ResetPassword(c Context) error {
	token := c.FormValue(ctrl.param.Token)
	password := c.FormValue(ctrl.param.Password)
	user := &inputdata.ResetPassword{
		Token:    token,
		Password: password,
	}
	err := ctrl.inputport.ResetPassword(user)
	if err != nil {
		c.String(statusCode(err), errs.Cause(err).Error())
		return err
	}
	return c.JSON(http.StatusOK, nil)
}

// Signout ...
// @summary Signout
// @description Delete cookie for login discrimination
// @tags UserAuth
// @produce json
// @param Authorization header string true "jwt token"
// @success 200
// @router /signout [post]
func (ctrl *UserController) Signout(c Context) error {
	return c.JSON(http.StatusOK, nil)
}

// generateHash ...
func (ctrl *UserController) generateHash(str string) (token string) {
	buf := make([]byte, 0)
	buf = append(buf, str...)
	buf = append(buf, os.Getenv("WS_PASSWORD_SALT")...)
	hash, err := bcrypt.GenerateFromPassword(buf, bcrypt.DefaultCost)
	if err != nil {
		return
	}
	token = string(hash)
	return
}
