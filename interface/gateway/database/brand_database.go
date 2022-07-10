package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/storage"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// BrandDatabase ...
type BrandDatabase struct {
	SQLHandler
}

// NewBrandDatabase ...
func NewBrandDatabase(sqlHandler SQLHandler) repository.BrandRepository {
	return &BrandDatabase{sqlHandler}
}

// FindListByUserID ...
func (db *BrandDatabase) FindListByUserID(userID string) ([]model.Brand, error) {
	brandsTable := &brandDB.BrandsTable
	statement := db.findStatement(brandsTable.UserID)
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{"statement": statement}).Error(err)
		return nil, err
	}
	defer rows.Close()
	brandList := []model.Brand{}
	for rows.Next() {
		brand := model.Brand{}
		scanList := generateScanList(&brand)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		brandList = append(brandList, brand)
	}
	return brandList, nil
}

// FindByID ...
func (db *BrandDatabase) FindByID(id string) (*model.Brand, error) {
	brandsTable := &brandDB.BrandsTable
	statement := db.findStatement(brandsTable.ID)
	valueList := []interface{}{id}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The brand doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	brand := &model.Brand{}
	scanList := generateScanList(brand)
	err = row.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	return brand, nil
}

// FindListByLikeUserID ...
func (db *BrandDatabase) FindListByLikeUserID(userID string) ([]model.Brand, error) {
	brandsTable := &brandDB.BrandsTable
	brandLikeTable := &brandDB.BrandLikeTable
	brandsActiveTable := &brandDB.BrandsActiveTable
	userInfoTable := &userDB.ProfileTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	statement := NewSQLBuilder().Select(brandsTable).
		RightJoin(brandLikeTable, brandLikeTable.BrandID, brandsTable.ID).
		RightJoin(brandsActiveTable, brandsActiveTable.BrandID, brandsTable.ID).
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, brandsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		LeftJoinWithColumns(mCategoriesTable, mCategoriesTable.ID, brandsTable.CategoryID).
		Match(brandLikeTable.UserID).
		Statement()
	valueList := []interface{}{userID}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer rows.Close()
	brandList := []model.Brand{}
	for rows.Next() {
		brand := model.Brand{}
		scanList := generateScanList(brand)
		rows.Scan(scanList...)
		brandList = append(brandList, brand)
	}
	return brandList, nil
}

// Save ...
func (db *BrandDatabase) Save(brand *model.Brand) error {
	_, err := db.Transaction(func() (interface{}, error) {
		brandsTable := &brandDB.BrandsTable
		statement := NewSQLBuilder().Insert(brandsTable)
		// Update image to storage
		err := storage.UploadFileForCreate(brand.LogoImage, storage.BrandPath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		valueList := generateValueList(brand)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		// Insert into active table
		brandsActiveTable := &brandDB.BrandsActiveTable
		statement = NewSQLBuilder().Insert(brandsActiveTable)
		valueList = []interface{}{brand.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Update ...
func (db *BrandDatabase) Update(brand *model.Brand) error {
	_, err := db.Transaction(func() (interface{}, error) {
		brandsTable := &brandDB.BrandsTable
		statement := NewSQLBuilder().Update(brandsTable)
		// Update image to storage
		err := storage.UploadFileForUpdate(brand.LogoImage, storage.BrandPath)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		valueList := generateValueList(brand)
		valueList = append(valueList, brand.ID)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *BrandDatabase) Delete(brand *model.Brand) error {
	_, err := db.Transaction(func() (interface{}, error) {
		brandsActiveTable := &brandDB.BrandsActiveTable
		statement := NewSQLBuilder().Delete(brandsActiveTable).Match(brandsActiveTable.BrandID).Statement()
		valueList := []interface{}{brand.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		brandsDeletedTable := &brandDB.BrandsDeletedTable
		statement = NewSQLBuilder().Insert(brandsDeletedTable)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// brandsTableQuery
func (db *BrandDatabase) findStatement(condition string) string {
	brandsTable := &brandDB.BrandsTable
	brandsActiveTable := &brandDB.BrandsActiveTable
	userInfoTable := &userDB.ProfileTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	statement := NewSQLBuilder().Select(brandsTable).
		RightJoin(brandsActiveTable, brandsActiveTable.BrandID, brandsTable.ID).
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, brandsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		LeftJoinWithColumns(mCategoriesTable, mCategoriesTable.ID, brandsTable.CategoryID).
		Match(condition).
		Statement()
	return statement
}
