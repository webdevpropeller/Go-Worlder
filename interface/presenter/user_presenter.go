package presenter

import (
	"go_worlder_system/consts"
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	sub = "sub"
	iat = "iat"
	exp = "exp"
)

// UserPresenter ...
type UserPresenter struct {
}

// NewUserPresenter ...
func NewUserPresenter() *UserPresenter {
	return &UserPresenter{}
}

// SignUp ...
func (presenter *UserPresenter) SignUp(user *model.User, token string) *outputdata.SignUp {
	url := strings.Join([]string{URL.Activate, "?", "token=", token}, "")
	return &outputdata.SignUp{
		Email:   user.Email,
		Message: url,
	}
}

// SignIn ...
func (presenter *UserPresenter) SignIn(user *model.User) *outputdata.SignIn {
	var oProfile *outputdata.Profile
	if user.Profile != nil {
		oProfile = &outputdata.Profile{
			Activity:   user.Profile.Activity,
			Industry:   user.Profile.Industry,
			Company:    user.Profile.Company,
			Country:    user.Profile.Country,
			Address1:   user.Profile.Address1,
			Address2:   user.Profile.Address2,
			URL:        user.Profile.URL,
			Phone:      user.Profile.Phone,
			Logo:       user.Profile.LogoPath,
			FirstName:  user.Profile.FirstName,
			MiddleName: user.Profile.MiddleName,
			FamilyName: user.Profile.FamilyName,
			Icon:       user.Profile.IconPath,
		}
	}
	// User
	oUser := outputdata.User{
		ID:      user.ID,
		Email:   user.Email,
		Profile: oProfile,
	}
	// JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		sub: user.ID,
		exp: time.Now().AddDate(0, 0, 7).Unix(),
		iat: time.Now().Unix(),
	})
	tokenString, _ := token.SignedString([]byte(os.Getenv(consts.JWT_SECRET)))
	return &outputdata.SignIn{
		JWT:  tokenString,
		User: oUser,
	}
}

// ForgotPassword ...
func (presenter *UserPresenter) ForgotPassword(user *model.User, token string) *outputdata.ForgotPassword {
	url := strings.Join([]string{URL.PasswordReset, "?", "token=", token}, "")
	return &outputdata.ForgotPassword{
		Email:   user.Email,
		Message: url,
	}
}
