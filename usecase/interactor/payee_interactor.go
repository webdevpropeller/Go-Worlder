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

// PayeeInteractor ...
type PayeeInteractor struct {
	outputport      outputport.PayeeOutputPort
	userRepository  repository.UserRepository
	payeeRepository repository.PayeeRepository
}

// NewPayeeInteractor ...
func NewPayeeInteractor(
	outputport outputport.PayeeOutputPort,
	userRepository repository.UserRepository,
	payeeRepository repository.PayeeRepository,
) (payeeInteractor *PayeeInteractor) {
	payeeInteractor = &PayeeInteractor{
		outputport:      outputport,
		userRepository:  userRepository,
		payeeRepository: payeeRepository,
	}
	return
}

// Index ...
func (interactor *PayeeInteractor) Index(userID string) ([]outputdata.Payee, error) {
	payeeList, err := interactor.payeeRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(payeeList), nil
}

// Create ...
func (interactor *PayeeInteractor) Create(iNewPayee *inputdata.NewPayee) error {
	user, err := interactor.userRepository.FindByID(iNewPayee.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	payee, err := model.NewPayee(user, iNewPayee.Name)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.payeeRepository.Save(payee)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Update ...
func (interactor *PayeeInteractor) Update(iUpdatedPayee *inputdata.UpdatedPayee) error {
	payee, err := interactor.payeeRepository.FindByID(iUpdatedPayee.ID)
	if err != nil {
		log.Error(err)
		return nil
	}
	if !payee.IsOwner(iUpdatedPayee.UserID) {
		errMsg := "The user can't update the payee"
		log.WithFields(log.Fields{
			idKey:     iUpdatedPayee.ID,
			userIDKey: iUpdatedPayee.UserID,
		}).Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = interactor.payeeRepository.Update(payee)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *PayeeInteractor) Delete(id string, userID string) error {
	payee, err := interactor.payeeRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !payee.IsOwner(userID) {
		errMsg := "The user can't delete the payee"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = interactor.payeeRepository.Delete(payee)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
