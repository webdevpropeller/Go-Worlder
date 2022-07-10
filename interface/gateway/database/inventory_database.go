package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"
	"strings"

	log "github.com/sirupsen/logrus"
)

// InventoryDatabase ...
type InventoryDatabase struct {
	SQLHandler
}

// NewInventoryDatabase ...
func NewInventoryDatabase(sqlHandler SQLHandler) repository.InventoryRepository {
	return &InventoryDatabase{sqlHandler}
}

// FindListByUserIDAndTypeID ...
func (db *InventoryDatabase) FindListByUserIDAndTypeID(userID string, inventoryTypeID int) ([]model.Inventory, error) {
	inventoryTable := &inventoryDB.InventoryTable
	userInfoTable := &userDB.ProfileTable
	productsTable := &productDB.ProductsTable
	productsActiveTable := &productDB.ProductsActiveTable
	brandsTable := &brandDB.BrandsTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	productInventoryTable := &productDB.ProductInventoryTable
	productManagementTable := &productDB.ProductManagementTable
	statement := NewSQLBuilder().Select(inventoryTable).
		// user
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, inventoryTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		// product
		LeftJoinWithColumns(productsTable, productsTable.ID, inventoryTable.ProductID).
		RightJoin(productsActiveTable, productsTable.ID, productsActiveTable.ProductID).
		LeftJoinWithColumns(brandsTable, productsTable.BrandID, brandsTable.ID).
		LeftJoinWithColumns(customeUserTable, brandsTable.UserID, customeUserTable.UserID).
		Columns(userInfoTable).
		LeftJoinWithColumns(mCategoriesTable, brandsTable.CategoryID, mCategoriesTable.ID).
		Columns(mCategoriesTable).
		LeftJoinWithColumns(productInventoryTable, productsTable.ID, productInventoryTable.ProductID).
		LeftJoinWithColumns(productManagementTable, productsTable.ID, productManagementTable.ProductID).
		Match(inventoryTable.UserID).
		Match(inventoryTable.InventoryTypeID).
		Statement()
	rows, err := db.Query(statement, userID, inventoryTypeID)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	inventoryList := []model.Inventory{}
	for rows.Next() {
		inventory := model.Inventory{}
		scanList := generateScanList(inventory)
		err = rows.Scan(scanList...)
		inventoryList = append(inventoryList, inventory)
	}
	return inventoryList, nil
}

// Create ...
func (db *InventoryDatabase) Create(inventory *model.Inventory) error {
	_, err := db.Transaction(func() (interface{}, error) {
		inventoryTable := inventoryDB.InventoryTable
		statement := NewSQLBuilder().Insert(inventoryTable)
		valueList := generateValueList(inventory)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, errs.Failed.Wrap(err, err.Error())
		}
		return nil, nil
	})
	return err
}

// Update ...
func (db *InventoryDatabase) Update(inventory *model.Inventory) error {
	_, err := db.Transaction(func() (interface{}, error) {
		inventoryTable := inventoryDB.InventoryTable
		statement := NewSQLBuilder().Update(inventoryTable)
		valueList := generateValueList(inventory)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, errs.Failed.Wrap(err, err.Error())
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *InventoryDatabase) Delete(inventory *model.Inventory) error {
	_, err := db.Transaction(func() (interface{}, error) {
		inventoryTable := inventoryDB.InventoryTable
		statement := NewSQLBuilder().Delete(inventoryTable).
			Match(inventoryTable.ID).
			Statement()
		valueList := []interface{}{inventory.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, errs.Failed.Wrap(err, err.Error())
		}
		return nil, nil
	})
	return err
}

func (db *InventoryDatabase) generateCommonStockQuery() string {
	inventoryColumn := &inventoryDB.InventoryTable
	return strings.Join([]string{
		"SUM(",
		rw.Case,
		rw.When, inventoryColumn.InventoryTypeID, "= 1", rw.Then, inventoryColumn.Quantity,
		rw.When, inventoryColumn.InventoryTypeID, "IN(2,3)", rw.Then, inventoryColumn.Quantity, "* -1",
		rw.Else, "0",
		rw.End,
		")",
	}, " ")
}
