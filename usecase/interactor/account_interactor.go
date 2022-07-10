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

// AccountInteractor ...
type AccountInteractor struct {
	outputport            outputport.AccountOutputPort
	userRepository        repository.UserRepository
	accountRepository     repository.AccountRepository
	accountTypeRepository repository.AccountTypeRepository
	accountNameRepository repository.AccountNameRepository
}

// NewAccountInteractor ...
func NewAccountInteractor(
	outputport outputport.AccountOutputPort,
	userRepository repository.UserRepository,
	accountRepository repository.AccountRepository,
	accountTypeRepository repository.AccountTypeRepository,
	accountNameRepository repository.AccountNameRepository,
) *AccountInteractor {
	return &AccountInteractor{
		outputport:            outputport,
		userRepository:        userRepository,
		accountRepository:     accountRepository,
		accountTypeRepository: accountTypeRepository,
		accountNameRepository: accountNameRepository,
	}
}

// Index ...
func (interactor *AccountInteractor) Index(userID string) ([]outputdata.Account, error) {
	accountList, err := interactor.accountRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(accountList), nil
}

// Show ...
func (interactor *AccountInteractor) Show(id string, userID string) (*outputdata.Account, error) {
	account, err := interactor.accountRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !account.IsOwner(userID) {
		errMsg := "The user can't get the account"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return interactor.outputport.Show(account), nil
}

// New ...
func (interactor *AccountInteractor) New() ([]outputdata.AccountName, error) {
	accountNameList, err := interactor.accountNameRepository.FindList()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.New(accountNameList), nil
}

// Create ...
func (interactor *AccountInteractor) Create(iAccount *inputdata.Account) error {
	accountName, err := interactor.accountNameRepository.FindByID(iAccount.AccountNameID)
	if err != nil {
		log.Error(err)
		return nil
	}
	user, err := interactor.userRepository.FindByID(iAccount.UserID)
	if err != nil {
		log.Error(err)
		return nil
	}
	account, err := model.NewAccount(user, accountName)
	if err != nil {
		log.Error(err)
		return nil
	}
	err = interactor.accountRepository.Save(account)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *AccountInteractor) Delete(id string, userID string) error {
	account, err := interactor.accountRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !account.IsOwner(userID) {
		msg, _ := log.WithFields(log.Fields{
			msgKey:    "The user can't delete the account",
			idKey:     account.ID,
			userIDKey: userID,
		}).String()
		return errs.Forbidden.New(msg)
	}
	err = interactor.accountRepository.Delete(account)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
