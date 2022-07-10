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

// TransactionInteractor ...
type TransactionInteractor struct {
	outputport            outputport.TransactionOutputPort
	userRepository        repository.UserRepository
	transactionRepository repository.TransactionRepository
	accountRepository     repository.AccountRepository
}

// NewTransactionInteractor ...
func NewTransactionInteractor(
	outputport outputport.TransactionOutputPort,
	userRepository repository.UserRepository,
	transactionRepository repository.TransactionRepository,
	accountRepository repository.AccountRepository,
) (transactionInteractor *TransactionInteractor) {
	transactionInteractor = &TransactionInteractor{
		outputport:            outputport,
		userRepository:        userRepository,
		transactionRepository: transactionRepository,
		accountRepository:     accountRepository,
	}
	return
}

// New ...
func (interactor *TransactionInteractor) New() ([]model.AccountItem, error) {
	accountItemList, err := interactor.transactionRepository.FindAccountItemList()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return accountItemList, nil
}

// Index ...
func (interactor *TransactionInteractor) Index(userID string) ([]outputdata.Transaction, error) {
	transactionList, err := interactor.transactionRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(transactionList), nil
}

// IndexAccount ...
func (interactor *TransactionInteractor) IndexAccount(accountID string, userID string) ([]outputdata.Transaction, error) {
	account, err := interactor.accountRepository.FindByID(accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !account.IsOwner(userID) {
		errMsg := "The user can't access the account"
		log.WithFields(log.Fields{
			idKey:     accountID,
			userIDKey: userID,
		})
		return nil, errs.Forbidden.New(errMsg)
	}
	transactionList, err := interactor.transactionRepository.FindListByAccountID(accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(transactionList), nil
}

// Create ...
func (interactor *TransactionInteractor) Create(iNewTransaction *inputdata.NewTransaction) error {
	user, err := interactor.userRepository.FindByID(iNewTransaction.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	accountItem, err := interactor.transactionRepository.FindAccountItemByID(iNewTransaction.AccountItemID)
	if err != nil {
		log.Error(err)
		return err
	}
	partner, err := interactor.transactionRepository.FindPartnerByID(iNewTransaction.PartnerID)
	if err != nil {
		log.Error(err)
		return err
	}
	var transactionAccountOption model.TransactionAccountOption
	if iNewTransaction.AccountID == "cash" {
		cash, err := interactor.accountRepository.FindCashByUserID(iNewTransaction.UserID)
		if err != nil {
			log.Error(err)
			return err
		}
		transactionAccountOption = model.TransactionByCash(cash)
	} else {
		account, err := interactor.accountRepository.FindByID(iNewTransaction.AccountID)
		if err != nil {
			log.Error(err)
			return err
		}
		if !account.IsOwner(iNewTransaction.UserID) {
			errMsg := "The user is not account owner"
			log.Error(errMsg)
			return errs.Forbidden.New(errMsg)
		}
		transactionAccountOption = model.TransactionByAccount(account)
	}
	transaction, err := model.NewTransaction(
		user, accountItem, partner, transactionAccountOption, iNewTransaction.IsIncome, iNewTransaction.IsPaid,
		iNewTransaction.AccrualDate, iNewTransaction.Amount, iNewTransaction.Remarks,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.transactionRepository.Save(transaction)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Edit ...
func (interactor *TransactionInteractor) Edit(id string, userID string) (*outputdata.Transaction, error) {
	transaction, err := interactor.transactionRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !transaction.IsOwner(userID) {
		msg, _ := log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).String()
		return nil, errs.Forbidden.New(msg)
	}
	return interactor.outputport.Edit(transaction), nil
}

// Update ...
func (interactor *TransactionInteractor) Update(iUpdatedTransaction *inputdata.UpdatedTransaction) error {
	transaction, err := interactor.transactionRepository.FindByID(iUpdatedTransaction.ID)
	if err != nil {
		log.Error(err)
		return err
	}
	accountItem, err := interactor.transactionRepository.FindAccountItemByID(iUpdatedTransaction.AccountItemID)
	if err != nil {
		log.Error(err)
		return err
	}
	partner, err := interactor.transactionRepository.FindPartnerByID(iUpdatedTransaction.PartnerID)
	if err != nil {
		log.Error(err)
		return err
	}
	var transactionAccountOption model.TransactionAccountOption
	if iUpdatedTransaction.AccountID == "cash" {
		cash, err := interactor.accountRepository.FindCashByUserID(iUpdatedTransaction.UserID)
		if err != nil {
			log.Error(err)
			return err
		}
		transactionAccountOption = model.TransactionByCash(cash)
	} else {
		account, err := interactor.accountRepository.FindByID(iUpdatedTransaction.AccountID)
		if err != nil {
			log.Error(err)
			return err
		}
		if !account.IsOwner(iUpdatedTransaction.UserID) {
			errMsg := "The user is not account owner"
			log.Error(errMsg)
			return errs.Forbidden.New(errMsg)
		}
		transactionAccountOption = model.TransactionByAccount(account)
	}
	transaction.Update(
		iUpdatedTransaction.UserID, accountItem, partner, transactionAccountOption, iUpdatedTransaction.IsIncome, iUpdatedTransaction.IsPaid,
		iUpdatedTransaction.AccrualDate, iUpdatedTransaction.Amount, iUpdatedTransaction.Remarks,
	)
	err = interactor.transactionRepository.Update(transaction)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *TransactionInteractor) Delete(id string, userID string) error {
	transaction, err := interactor.transactionRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !transaction.IsOwner(userID) {
		errMsg := "The user can't update the transaction"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error(errMsg)
		return errs.Forbidden.Errorf(errMsg)
	}
	err = interactor.transactionRepository.Delete(transaction)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
