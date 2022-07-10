package interactor

import (
	"go_worlder_system/domain/model"
	inputdata "go_worlder_system/usecase/input/data"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// OrderInteractor ...
type OrderInteractor struct {
	userRepository    repository.UserRepository
	orderRepository   repository.OrderRepository
	productRepository repository.ProductRepository
}

// NewOrderInteractor ...
func NewOrderInteractor(
	userRepository repository.UserRepository,
	orderRepository repository.OrderRepository,
	productRepository repository.ProductRepository,
) (orderInteracotr *OrderInteractor) {
	orderInteracotr = &OrderInteractor{
		userRepository:    userRepository,
		orderRepository:   orderRepository,
		productRepository: productRepository,
	}
	return
}

// Create ...
func (interactor *OrderInteractor) Create(iOrder *inputdata.Order) (err error) {
	user, err := interactor.userRepository.FindByID(iOrder.UserID)
	if err != nil {
		log.Error(err)
		return err
	}
	for _, purchase := range iOrder.PurchaseList {
		product, err := interactor.productRepository.FindByID(purchase.ProductID)
		if err != nil {
			log.Error(err)
			return err
		}
		order, err := model.NewOrder(user, product, purchase.Quantity)
		if err != nil {
			log.Error(err)
			return err
		}
		err = interactor.orderRepository.Create(order)
		if err != nil {
			log.Error(err)
			return err
		}
	}
	return
}
