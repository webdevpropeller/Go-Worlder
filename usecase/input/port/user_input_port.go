package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// UserInputPort ...
type UserInputPort interface {
	SignUp(*inputdata.SignUp) (*outputdata.SignUp, error)
	Activate(token string) error
	SignIn(*inputdata.SignIn) (*outputdata.SignIn, error)
	ForgotPassword(*inputdata.ForgotPassword) (*outputdata.ForgotPassword, error)
	ResetPassword(*inputdata.ResetPassword) error
}
