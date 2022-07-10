package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/storage"
	"go_worlder_system/usecase/repository"
	"strings"

	log "github.com/sirupsen/logrus"
)

// ProductDatabase ...
type ProductDatabase struct {
	SQLHandler
}

// NewProductDatabase ...
func NewProductDatabase(sqlHandler SQLHandler) repository.ProductRepository {
	return &ProductDatabase{sqlHandler}
}

// RetrieveListByKeyWord ...
func (db *ProductDatabase) RetrieveListByKeyWord(keyword string) ([]model.Product, error) {
	productsTable := &productDB.ProductsTable
	productsActiveTable := &productDB.ProductsActiveTable
	brandsTable := &brandDB.BrandsTable
	userInfoTable := &userDB.ProfileTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	productInventoryTable := &productDB.ProductInventoryTable
	productManagementTable := &productDB.ProductManagementTable
	statement := NewSQLBuilder().Select(productsTable).
		RightJoin(productsActiveTable, productsTable.ID, productsActiveTable.ProductID).
		LeftJoinWithColumns(brandsTable, productsTable.BrandID, brandsTable.ID).
		LeftJoinWithColumns(customeUserTable, brandsTable.UserID, customeUserTable.UserID).
		LeftJoinWithColumns(userInfoTable, brandsTable.UserID, userInfoTable.UserID).
		LeftJoinWithColumns(mCategoriesTable, brandsTable.CategoryID, mCategoriesTable.ID).
		Columns(mCategoriesTable).
		LeftJoinWithColumns(productInventoryTable, productsTable.ID, productInventoryTable.ProductID).
		LeftJoinWithColumns(productManagementTable, productsTable.ID, productManagementTable.ProductID).
		Like(productsTable.Name).
		Statement()
	keyword = strings.Join([]string{"%", keyword, "%"}, "")
	rows, err := db.Query(statement, keyword)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	productList := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		scanList := generateScanList(&product)
		rows.Scan(scanList...)
		productList = append(productList, product)
	}
	return productList, nil
}

// FindListByUserID ...
func (db *ProductDatabase) FindListByUserID(userID string) ([]model.Product, error) {
	brandsTable := &brandDB.BrandsTable
	statement := db.findStatement(brandsTable.UserID)
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	productList := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		scanList := generateScanList(&product)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		productList = append(productList, product)
	}
	return productList, nil
}

// FindListByBrandID ...
func (db *ProductDatabase) FindListByBrandID(brandID string) ([]model.Product, error) {
	productsTable := &productDB.ProductsTable
	statement := db.findStatement(productsTable.ID)
	valueList := []interface{}{brandID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	productList := []model.Product{}
	for rows.Next() {
		product := model.Product{}
		scanList := generateScanList(&product)
		rows.Scan(scanList...)
		productList = append(productList, product)
	}
	return productList, nil
}

// FindByID ...
func (db *ProductDatabase) FindByID(id string) (*model.Product, error) {
	productsTable := &productDB.ProductsTable
	statement := db.findStatement(productsTable.ID)
	valueList := []interface{}{id}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The product doesn'T exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	product := &model.Product{}
	scanList := generateScanList(product)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return product, nil
}

// Create ...
func (db *ProductDatabase) Create(product *model.Product) error {
	_, err := db.Transaction(func() (interface{}, error) {
		productsTable := &productDB.ProductsTable
		statement := NewSQLBuilder().Insert(productsTable)
		// Update image to storage
		err := storage.UploadFileForCreate(product.Image, storage.ProductPath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		valueList := generateValueList(product)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into active table
		productsActiveTable := &productDB.ProductsActiveTable
		statement = NewSQLBuilder().Insert(productsActiveTable)
		valueList = []interface{}{product.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into inventory table
		productInventoryTable := &productDB.ProductInventoryTable
		statement = NewSQLBuilder().Insert(productInventoryTable)
		valueList = generateValueList(product.Inventory)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// insert into management table
		productManagementTable := &productDB.ProductManagementTable
		statement = NewSQLBuilder().Insert(productManagementTable)
		valueList = generateValueList(product.Management)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Update ...
func (db *ProductDatabase) Update(product *model.Product) error {
	_, err := db.Transaction(func() (interface{}, error) {
		productsTable := &productDB.ProductsTable
		statement := NewSQLBuilder().Update(productsTable)
		// Update image to storage
		err := storage.UploadFileForUpdate(product.Image, storage.ProductPath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		valueList := generateValueList(product)
		valueList = append(valueList, product.ID)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// UpdateInventory ...
func (db *ProductDatabase) UpdateInventory(inventory *model.ProductInventory) error {
	productInventoryTable := &productDB.ProductInventoryTable
	statement := NewSQLBuilder().Update(productInventoryTable)
	valueList := generateValueList(inventory)
	valueList = append(valueList, inventory.ProductID)
	_, err := db.Exec(statement, valueList...)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// UpdateManagement ...
func (db *ProductDatabase) UpdateManagement(manegement *model.ProductManagement) error {
	productManagementTable := &productDB.ProductManagementTable
	statement := NewSQLBuilder().Update(productManagementTable)
	valueList := generateValueList(manegement)
	valueList = append(valueList, manegement.ProductID)
	_, err := db.Exec(statement, valueList...)
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

// Delete ...
func (db *ProductDatabase) Delete(id string) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Delete from active table
		productsActiveTable := &productDB.ProductsActiveTable
		statement := NewSQLBuilder().Delete(productsActiveTable).Match(productsActiveTable.ProductID).Statement()
		valueList := []interface{}{id}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into deleted table
		productsDeletedTable := &productDB.ProductsDeletedTable
		statement = NewSQLBuilder().Insert(productsDeletedTable)
		valueList = []interface{}{id}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// DeleteListByBrandID ...
func (db *ProductDatabase) DeleteListByBrandID(brandID string) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Delete from active table
		productsActiveTable := &productDB.ProductsActiveTable
		productsTable := &productDB.ProductsTable
		statement := strings.Join([]string{
			rw.Delete, productsActiveTable.NAME(),
			rw.From, productsActiveTable.NAME(),
			rw.LeftJoin, productsTable.Name,
			rw.On, productsActiveTable.ProductID, "=", productsTable.ID,
			rw.Where, productsTable.BrandID, "= ?",
		}, " ")
		log.Println(statement)
		valueList := []interface{}{brandID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Insert into deleted table
		productsDeletedTable := &productDB.ProductsDeletedTable
		columnList := generateColumnList(productsDeletedTable)
		columnQuery := generateInsertColumnQuery(columnList)
		statement = strings.Join([]string{
			rw.InsertInto, productsDeletedTable.NAME(), columnQuery,
			rw.Select, productsTable.ID,
			rw.From, productsTable.NAME(),
			rw.Where, productsTable.BrandID, "= ?",
		}, " ")
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// getProductsTableQuery
func (db *ProductDatabase) findStatement(condition string) string {
	productsTable := &productDB.ProductsTable
	productsActiveTable := &productDB.ProductsActiveTable
	brandsTable := &brandDB.BrandsTable
	userInfoTable := &userDB.ProfileTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	productInventoryTable := &productDB.ProductInventoryTable
	productManagementTable := &productDB.ProductManagementTable
	statement := NewSQLBuilder().Select(productsTable).
		RightJoin(productsActiveTable, productsTable.ID, productsActiveTable.ProductID).
		LeftJoinWithColumns(brandsTable, productsTable.BrandID, brandsTable.ID).
		LeftJoinWithColumns(customeUserTable, brandsTable.UserID, customeUserTable.UserID).
		LeftJoinWithColumns(userInfoTable, brandsTable.UserID, userInfoTable.UserID).
		LeftJoinWithColumns(mCategoriesTable, brandsTable.CategoryID, mCategoriesTable.ID).
		Columns(mCategoriesTable).
		LeftJoinWithColumns(productInventoryTable, productsTable.ID, productInventoryTable.ProductID).
		LeftJoinWithColumns(productManagementTable, productsTable.ID, productManagementTable.ProductID).
		Match(condition).
		Statement()
	return statement
}
