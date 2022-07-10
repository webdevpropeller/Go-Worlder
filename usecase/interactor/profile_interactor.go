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

// ProfileInteractor ...
type ProfileInteractor struct {
	outputport        outputport.ProfileOutputPort
	userRepository    repository.UserRepository
	optionRepository  repository.OptionRespository
	brandRepository   repository.BrandRepository
	productRepository repository.ProductRepository
	projectRepository repository.ProjectRepository
}

// NewProfileInteractor ...
func NewProfileInteractor(
	profileOutputPort outputport.ProfileOutputPort,
	userRepository repository.UserRepository,
	optionRepository repository.OptionRespository,
	brandRepository repository.BrandRepository,
	productRepository repository.ProductRepository,
	projectRepository repository.ProjectRepository,
) *ProfileInteractor {
	return &ProfileInteractor{
		outputport:        profileOutputPort,
		userRepository:    userRepository,
		optionRepository:  optionRepository,
		brandRepository:   brandRepository,
		productRepository: productRepository,
		projectRepository: projectRepository,
	}
}

// Show ...
func (interactor *ProfileInteractor) Show(accountID string) (*outputdata.PublicUser, error) {
	profile, err := interactor.userRepository.FindByAccountID(accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Show(profile), nil
}

func (interactor *ProfileInteractor) New() (*outputdata.ProfileSelectItem, error) {
	industries, err := interactor.optionRepository.FindIndustries()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	countries, err := interactor.optionRepository.FindCountries()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	cardCompanies, err := interactor.optionRepository.FindCardCompanies()
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.New(industries, countries, cardCompanies), nil
}

// Create ...
func (interactor *ProfileInteractor) Create(iProfile *inputdata.Profile) (*outputdata.User, error) {
	user, err := interactor.userRepository.FindByID(iProfile.UserID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	if user.Profile != nil {
		errMsg := "User info already exists"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.Conflict.New(errMsg)
	}
	card, err := model.NewCard(
		user,
		iProfile.Card.Company,
		iProfile.Card.Name,
		iProfile.Card.Number,
		iProfile.Card.Expiry,
		iProfile.Card.SecurityCode,
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.userRepository.SaveCard(card)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = user.CreateProfile(
		iProfile.ActivityID,
		iProfile.IndustryID,
		iProfile.Company,
		iProfile.CountryID,
		iProfile.Address1,
		iProfile.Address2,
		iProfile.ZipCode,
		iProfile.URL,
		iProfile.Phone,
		iProfile.AccountID,
		iProfile.Logo,
		iProfile.FirstName,
		iProfile.MiddleName,
		iProfile.FamilyName,
		iProfile.Icon,
	)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	err = interactor.userRepository.SaveProfile(user)
	if err != nil {
		log.Error(err)
		interactor.userRepository.DeleteCard(card)
		return nil, err
	}
	return interactor.outputport.Create(user), nil
}

// Edit ...
func (interactor *ProfileInteractor) Edit(userID string) (*outputdata.Profile, error) {
	user, err := interactor.userRepository.FindByID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Edit(user.Profile), nil
}

// Update ...
func (interactor *ProfileInteractor) Update(iProfile *inputdata.Profile) error {
	user, err := interactor.userRepository.FindByID(iProfile.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	err = user.Update(
		iProfile.ActivityID,
		iProfile.IndustryID,
		iProfile.Company,
		iProfile.CountryID,
		iProfile.Address1,
		iProfile.Address2,
		iProfile.ZipCode,
		iProfile.URL,
		iProfile.Phone,
		iProfile.AccountID,
		iProfile.Logo,
		iProfile.FirstName,
		iProfile.MiddleName,
		iProfile.FamilyName,
		iProfile.Icon,
	)
	if err != nil {
		log.Error(err)
		return err
	}
	err = interactor.userRepository.UpdateProfile(user)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// ShowProfile ...
func (interactor *ProfileInteractor) ShowProfile(userID string) (*outputdata.Profile, error) {
	_, err := interactor.userRepository.FindByID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	user, err := interactor.userRepository.FindByID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	brandList, err := interactor.brandRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	productList, err := interactor.productRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	projectList, err := interactor.projectRepository.FindListByUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.Profile(user.Profile, brandList, productList, projectList), nil
}

// IndexBrandLike ...
func (interactor *ProfileInteractor) IndexBrandLike(userID string) ([]outputdata.Brand, error) {
	brandList, err := interactor.brandRepository.FindListByLikeUserID(userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return interactor.outputport.BrandLike(brandList), nil
}
