package interactor

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// BrandLikeInteractor ...
type BrandLikeInteractor struct {
	outputport          outputport.BrandLikeOutputPort
	userRepository      repository.UserRepository
	brandRepository     repository.BrandRepository
	brandLikeRepository repository.BrandLikeRepository
}

// NewBrandLikeInteractor ...
func NewBrandLikeInteractor(
	outputport outputport.BrandLikeOutputPort,
	userRepository repository.UserRepository,
	brandRepository repository.BrandRepository,
	brandLikeRepository repository.BrandLikeRepository,
) *BrandLikeInteractor {
	return &BrandLikeInteractor{
		outputport:          outputport,
		userRepository:      userRepository,
		brandRepository:     brandRepository,
		brandLikeRepository: brandLikeRepository,
	}
}

// Index ...
func (interactor *BrandLikeInteractor) Index(brandID string) ([]outputdata.Profile, error) {
	userList, err := interactor.userRepository.FindListByLikeBrandID(brandID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Index(userList), nil
}

// Create ...
func (interactor *BrandLikeInteractor) Create(brandID string, userID string) error {
	brand, err := interactor.brandRepository.FindByID(brandID)
	if err != nil {
		log.Error(err)
		return err
	}
	user, err := interactor.userRepository.FindByID(userID)
	if err != nil {
		log.Error(err)
		return err
	}
	brandLike, err := model.NewBrandLike(brand, user)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.brandLikeRepository.Save(brandLike)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (interactor *BrandLikeInteractor) Delete(brandID string, userID string) error {
	brandLike, err := interactor.brandLikeRepository.FindByBrandIDAndUserID(brandID, userID)
	err = interactor.brandLikeRepository.Delete(brandLike)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}
