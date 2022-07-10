package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// UserOutputPort ...
type UserOutputPort interface {
	SignUp(user *model.User, token string) *outputdata.SignUp
	SignIn(*model.User) *outputdata.SignIn
	ForgotPassword(user *model.User, token string) *outputdata.ForgotPassword
}
