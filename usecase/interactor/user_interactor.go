package interactor

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// UserInteractor express the flow of authentication
type UserInteractor struct {
	outputport          outputport.UserOutputPort
	userRepository      repository.UserRepository
	userTokenRepository repository.UserTokenRepository
}

// NewUserInteractor ...
func NewUserInteractor(
	outputport outputport.UserOutputPort,
	userRepository repository.UserRepository,
	userTokenRepository repository.UserTokenRepository,
) *UserInteractor {
	return &UserInteractor{
		outputport:          outputport,
		userRepository:      userRepository,
		userTokenRepository: userTokenRepository,
	}
}

// SignUp ...
func (interactor *UserInteractor) SignUp(iSignUp *inputdata.SignUp) (*outputdata.SignUp, error) {
	existingUser, _ := interactor.userRepository.FindByEmail(iSignUp.User.Email)
	if existingUser != nil {
		errMsg := "The email already exists"
		log.Error(errMsg)
		return nil, errs.Conflict.New(errMsg)
	}
	user, err := model.NewUser(iSignUp.User.Email, iSignUp.User.Password, iSignUp.User.Category)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.userRepository.Save(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	userToken, err := model.NewUserToken(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.userTokenRepository.Save(userToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.SignUp(user, userToken.Token), nil
}

// Activate ...
func (interactor *UserInteractor) Activate(token string) error {
	userToken, err := interactor.userTokenRepository.FindByToken(token)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.userRepository.UpdateToActive(userToken.User)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.userTokenRepository.Delete(userToken)
	return nil
}

// SignIn ...
func (interactor *UserInteractor) SignIn(iSignIn *inputdata.SignIn) (*outputdata.SignIn, error) {
	user, err := interactor.userRepository.FindByEmail(iSignIn.Email)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !user.IsValidPassword(iSignIn.Password) {
		errMsg := "The password is invalid"
		log.Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return interactor.outputport.SignIn(user), nil
}

// ForgotPassword ...
func (interactor *UserInteractor) ForgotPassword(iForgotPassword *inputdata.ForgotPassword) (*outputdata.ForgotPassword, error) {
	user, err := interactor.userRepository.FindByEmail(iForgotPassword.Email)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	userToken, err := model.NewUserToken(user)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.userTokenRepository.Save(userToken)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.ForgotPassword(user, userToken.Token), nil
}

// ResetPassword ...
func (interactor *UserInteractor) ResetPassword(iResetPassword *inputdata.ResetPassword) error {
	userToken, err := interactor.userTokenRepository.FindByToken(iResetPassword.Token)
	if err != nil {
		log.Error(err)
		return err
	}
	user := userToken.User
	user.UpdatePassword(iResetPassword.Password)
	err = interactor.userRepository.UpdatePassword(user)
	if err != nil {
		log.Error(err)
		return err
	}
	interactor.userTokenRepository.Delete(userToken)
	return nil
}
