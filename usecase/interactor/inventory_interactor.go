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

// InventoryInteractor ...
type InventoryInteractor struct {
	outputport          outputport.InventoryOutputPort
	userRepository      repository.UserRepository
	inventoryRepository repository.InventoryRepository
	productRepository   repository.ProductRepository
}

// NewInventoryInteractor ...
func NewInventoryInteractor(
	outputport outputport.InventoryOutputPort,
	userRepository repository.UserRepository,
	inventoryRepository repository.InventoryRepository,
	productRepository repository.ProductRepository,
) (inventoryInteractor *InventoryInteractor) {
	inventoryInteractor = &InventoryInteractor{
		outputport:          outputport,
		userRepository:      userRepository,
		inventoryRepository: inventoryRepository,
		productRepository:   productRepository,
	}
	return
}

// IndexProduct ...
func (interactor *InventoryInteractor) IndexProduct(userID string) ([]outputdata.Inventory, error) {
	productList, err := interactor.productRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.IndexProduct(productList), nil
}

// ShowProduct ...
func (interactor *InventoryInteractor) ShowProduct(productID string, userID string) (*outputdata.Inventory, error) {
	product, err := interactor.productRepository.FindByID(productID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !product.IsOwner(userID) {
		return nil, errs.Forbidden.New("The user can't access the product")
	}
	return interactor.outputport.ShowProduct(product), nil
}

// EditProduct ...
func (interactor *InventoryInteractor) EditProduct(productID string, userID string) (*outputdata.Inventory, error) {
	product, err := interactor.productRepository.FindByID(productID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !product.IsOwner(userID) {
		errMsg := "The user can't get the product inventory"
		log.WithFields(log.Fields{
			idKey:     product.ID,
			userIDKey: userID,
		}).Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return interactor.outputport.Editproduct(product), nil
}

// UpdateProduct ...
func (interactor *InventoryInteractor) UpdateProduct(iManagement *inputdata.ProductManagement, userID string) error {
	product, err := interactor.productRepository.FindByID(iManagement.ProductID)
	if err != nil {
		log.Error(err)
		return err
	}
	if !product.IsOwner(userID) {
		errMsg := "The user can't update the product inventory"
		log.WithFields(log.Fields{
			idKey:     iManagement.ProductID,
			userIDKey: userID,
		}).Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = product.UpdateManagement(iManagement.Storage, iManagement.Barcode, iManagement.Memo)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.productRepository.UpdateManagement(product.Management)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// New ...
func (interactor *InventoryInteractor) New(userID string) ([]outputdata.Inventory, error) {
	productList, err := interactor.productRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.New(productList), nil
}

// Create ...
func (interactor *InventoryInteractor) Create(iInventory *inputdata.Inventory) error {
	user, err := interactor.userRepository.FindByID(iInventory.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, update := range iInventory.Request.UpdateList {
		product, err := interactor.productRepository.FindByID(update.ProductID)
		if err != nil {
			log.Error(err)
			return err
		}
		if !product.IsOwner(iInventory.UserID) {
			errMsg := "The user can't access the product"
			log.WithFields(log.Fields{
				idKey:     update.ProductID,
				userIDKey: iInventory.UserID,
			}).Error(errMsg)
			return errs.Forbidden.New(errMsg)
		}
		inventory, err := model.NewInventory(user, product, iInventory.Type, update.Quantity)
		if err != nil {
			log.Error(err)
			return err
		}
		err = interactor.inventoryRepository.Create(inventory)
		if err != nil {
			log.Error(err)
			return err
		}
		err = interactor.productRepository.UpdateInventory(product.Inventory)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}

// UpdateStocktaking ...
func (interactor *InventoryInteractor) UpdateStocktaking(iStocktaking *inputdata.Stocktaking) error {
	for _, update := range iStocktaking.Request.UpdateList {
		product, err := interactor.productRepository.FindByID(update.ProductID)
		if err != nil {
			log.Error(err)
			return nil
		}
		if !product.IsOwner(iStocktaking.UserID) {
			msg, _ := log.WithFields(log.Fields{
				msgKey:   "The user can't take stock",
				"update": update,
			}).String()
			return errs.Forbidden.New(msg)
		}
		err = product.Stocktaking(update.Stock)
		if err != nil {
			log.Error(err)
			return err
		}
		err = interactor.productRepository.UpdateInventory(product.Inventory)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return nil
}
