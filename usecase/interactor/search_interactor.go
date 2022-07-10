package interactor

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
	outputport "go_worlder_system/usecase/output/port"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// SearchInteractor ...
type SearchInteractor struct {
	outputport        outputport.SearchOutputPort
	productRepository repository.ProductRepository
}

// NewSearchInteractor ...
func NewSearchInteractor(
	outputport outputport.SearchOutputPort,
	productRepository repository.ProductRepository,
) *SearchInteractor {
	return &SearchInteractor{
		outputport:        outputport,
		productRepository: productRepository,
	}
}

// SearchProduct ...
func (interactor *SearchInteractor) SearchProduct(iSearch *inputdata.Search) ([]outputdata.Product, error) {
	productList, err := interactor.productRepository.RetrieveListByKeyWord(iSearch.Keyword)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return interactor.outputport.SearchProduct(productList), nil
}
