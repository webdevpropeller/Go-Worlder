package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// BrandLikeDatabase ...
type BrandLikeDatabase struct {
	SQLHandler
}

// NewBrandLikeDatabase ...
func NewBrandLikeDatabase(sqlHandler SQLHandler) repository.BrandLikeRepository {
	return &BrandLikeDatabase{sqlHandler}
}

// FindByBrandIDAndUserID ...
func (db *BrandLikeDatabase) FindByBrandIDAndUserID(brandID string, userID string) (*model.BrandLike, error) {
	statement := db.findStatement()
	valueList := []interface{}{brandID, userID}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The brand like doesn't exist"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	brandLike := &model.BrandLike{}
	scanList := generateScanList(brandLike)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return brandLike, nil
}

// Save ...
func (db *BrandLikeDatabase) Save(brandLike *model.BrandLike) error {
	_, err := db.Transaction(func() (interface{}, error) {
		brandLikeTable := &brandDB.BrandLikeTable
		statement := NewSQLBuilder().Insert(brandLikeTable)
		valueList := generateValueList(brandLike)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *BrandLikeDatabase) Delete(brandLike *model.BrandLike) error {
	_, err := db.Transaction(func() (interface{}, error) {
		brandLikeTable := &brandDB.BrandLikeTable
		statement := NewSQLBuilder().Delete(brandLikeTable).Match(brandLikeTable.BrandID).Match(brandLikeTable.UserID).Statement()
		valueList := generateValueList(brandLike)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (db *BrandLikeDatabase) findStatement() string {
	brandLikeTable := &brandDB.BrandLikeTable
	brandsTable := &brandDB.BrandsTable
	userInfoTable := &userDB.ProfileTable
	mCategoriesTable := &categoryDB.MCategoriesTable
	subCustomeUserTable := alias(customeUserTable, "sub").(CustomeUserTable)
	statement := NewSQLBuilder().Select(brandLikeTable).
		LeftJoinWithColumns(brandsTable, brandsTable.ID, brandLikeTable.BrandID).
		LeftJoinWithColumns(customeUserTable, customeUserTable.UserID, brandsTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		LeftJoinWithColumns(mCategoriesTable, mCategoriesTable.ID, brandsTable.CategoryID).
		LeftJoinWithColumns(subCustomeUserTable, subCustomeUserTable.UserID, subCustomeUserTable.UserID).
		LeftJoinWithColumns(userInfoTable, userInfoTable.UserID, customeUserTable.UserID).
		Match(brandLikeTable.BrandID).
		Match(brandLikeTable.UserID).
		Statement()
	return statement
}
