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

// ProductInteractor ...
type ProductInteractor struct {
	outputport         outputport.ProductOutputPort
	userRepository     repository.UserRepository
	productRepository  repository.ProductRepository
	brandRepository    repository.BrandRepository
	categoryRepository repository.CategoryRepository
}

// NewProductInteractor ...
func NewProductInteractor(
	outputport outputport.ProductOutputPort,
	userRepository repository.UserRepository,
	productRepository repository.ProductRepository,
	brandRepository repository.BrandRepository,
	categoryRepository repository.CategoryRepository,
) *ProductInteractor {
	return &ProductInteractor{
		outputport:         outputport,
		userRepository:     userRepository,
		productRepository:  productRepository,
		brandRepository:    brandRepository,
		categoryRepository: categoryRepository,
	}
}

// Index ...
func (interactor *ProductInteractor) Index(userID string) ([]outputdata.Product, error) {
	productList, err := interactor.productRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(productList), nil
}

// Show ...
func (interactor *ProductInteractor) Show(id string) (*outputdata.Product, error) {
	product, err := interactor.productRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Show(product), nil
}

// Edit ...
func (interactor *ProductInteractor) Edit(id string, userID string) (*outputdata.Product, error) {
	product, err := interactor.productRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !product.IsOwner(userID) {
		msg, _ := log.WithFields(log.Fields{
			msgKey: "The user can't get the product",
			idKey:  id,
		}).String()
		return nil, errs.Forbidden.New(msg)
	}
	return interactor.outputport.Edit(product), nil
}

// Create ...
func (interactor *ProductInteractor) Create(iNewProduct *inputdata.NewProduct) error {
	brand, err := interactor.brandRepository.FindByID(iNewProduct.BrandID)
	if err != nil {
		log.Error(err)
		return err
	}
	if !brand.IsOwner(iNewProduct.UserID) {
		errMsg := "The user is not owner of the brand"
		log.Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	category, err := interactor.categoryRepository.FindByID(iNewProduct.CategoryID)
	if err != nil {
		log.Error(err)
		return err
	}
	product, err := model.NewProduct(
		brand, category, iNewProduct.GenreID, iNewProduct.Name,
		iNewProduct.Price, iNewProduct.Image,
		iNewProduct.Description, iNewProduct.IsDraft,
	)
	err = interactor.productRepository.Create(product)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Update ...
func (interactor *ProductInteractor) Update(iUpdatedProduct *inputdata.UpdatedProduct) error {
	product, err := interactor.productRepository.FindByID(iUpdatedProduct.ID)
	if err != nil {
		log.Error(err)
		return err
	}
	if !product.IsOwner(iUpdatedProduct.UserID) {
		msg := "The user can't update the product"
		log.Error(msg)
		return errs.Forbidden.New(msg)
	}
	category, err := interactor.categoryRepository.FindByID(iUpdatedProduct.CategoryID)
	if err != nil {
		log.Error(err)
		return err
	}
	err = product.Update(
		category, iUpdatedProduct.GenreID, iUpdatedProduct.Name,
		iUpdatedProduct.Price, iUpdatedProduct.Image,
		iUpdatedProduct.Description, iUpdatedProduct.IsDraft,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.productRepository.Update(product)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *ProductInteractor) Delete(id string, userID string) error {
	product, err := interactor.productRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !product.IsOwner(userID) {
		errMsg := "The user can't delete the product"
		log.Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = interactor.productRepository.Delete(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
