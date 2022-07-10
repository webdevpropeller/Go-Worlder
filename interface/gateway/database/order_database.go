package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/usecase/repository"

	"github.com/labstack/gommon/log"
)

// OrderDatabase ...
type OrderDatabase struct {
	SQLHandler
}

// NewOrderDatabase ...
func NewOrderDatabase(sqlHandler SQLHandler) repository.OrderRepository {
	return &OrderDatabase{sqlHandler}
}

// Create ...
func (db *OrderDatabase) Create(order *model.Order) error {
	_, err := db.Transaction(func() (interface{}, error) {
		ordersTable := orderDB.OrdersTable
		statement := NewSQLBuilder().Insert(ordersTable)
		valueList := generateValueList(order)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
