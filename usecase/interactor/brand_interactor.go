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

// BrandInteractor ...
type BrandInteractor struct {
	outputport         outputport.BrandOutputPort
	userRepository     repository.UserRepository
	brandRepository    repository.BrandRepository
	productRepository  repository.ProductRepository
	categoryRepository repository.CategoryRepository
}

// NewBrandInteractor ...
func NewBrandInteractor(
	outputport outputport.BrandOutputPort,
	userRepository repository.UserRepository,
	brandRepository repository.BrandRepository,
	productRepository repository.ProductRepository,
	categoryRepository repository.CategoryRepository,
) *BrandInteractor {
	return &BrandInteractor{
		outputport:         outputport,
		userRepository:     userRepository,
		brandRepository:    brandRepository,
		productRepository:  productRepository,
		categoryRepository: categoryRepository,
	}
}

// Index ...
func (interactor *BrandInteractor) Index(userID string) ([]outputdata.Brand, error) {
	brandList, err := interactor.brandRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(brandList), nil
}

// Show ...
func (interactor *BrandInteractor) Show(id string) (*outputdata.Brand, error) {
	brand, err := interactor.brandRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Show(brand), nil
}

// Create ...
func (interactor *BrandInteractor) Create(iNewBrand *inputdata.NewBrand) error {
	user, err := interactor.userRepository.FindByID(iNewBrand.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	category, err := interactor.categoryRepository.FindByID(iNewBrand.CategoryID)
	if err != nil {
		log.Error(err)
		return err
	}
	brand, err := model.NewBrand(
		user,
		category,
		iNewBrand.Name,
		iNewBrand.Slogan,
		iNewBrand.LogoImage,
		iNewBrand.Description,
		iNewBrand.IsDraft,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.brandRepository.Save(brand)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Edit ...
func (interactor *BrandInteractor) Edit(id string, userID string) (*outputdata.Brand, error) {
	brand, err := interactor.brandRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if !brand.IsOwner(userID) {
		errMsg := "The user can't get the brand"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error(errMsg)
		return nil, errs.Forbidden.New(errMsg)
	}
	return interactor.outputport.Edit(brand), nil
}

// Update ...
func (interactor *BrandInteractor) Update(iUpdatedBrand *inputdata.UpdatedBrand) error {
	brand, err := interactor.brandRepository.FindByID(iUpdatedBrand.ID)
	if err != nil {
		log.Error(err)
		return err
	}
	if !brand.IsOwner(iUpdatedBrand.UserID) {
		errMsg := "The user can't get the brand"
		log.WithFields(log.Fields{}).Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	category, err := interactor.categoryRepository.FindByID(iUpdatedBrand.CategoryID)
	if err != nil {
		log.Error(err)
		return err
	}
	err = brand.Update(
		category,
		iUpdatedBrand.Name,
		iUpdatedBrand.Slogan,
		iUpdatedBrand.LogoImage,
		iUpdatedBrand.Description,
		iUpdatedBrand.IsDraft,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.brandRepository.Update(brand)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *BrandInteractor) Delete(id string, userID string) error {
	brand, err := interactor.brandRepository.FindByID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	if !brand.IsOwner(userID) {
		errMsg := "The user can't delete the brand"
		log.WithFields(log.Fields{
			idKey:     id,
			userIDKey: userID,
		}).Error(errMsg)
		return errs.Forbidden.New(errMsg)
	}
	err = interactor.brandRepository.Delete(brand)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.productRepository.DeleteListByBrandID(id)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
