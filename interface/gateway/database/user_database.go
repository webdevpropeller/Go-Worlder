package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/storage"
	"strings"

	log "github.com/sirupsen/logrus"
)

// UserDatabase ...
type UserDatabase struct {
	SQLHandler
}

// NewUserDatabase ...
func NewUserDatabase(sqlHandler SQLHandler) *UserDatabase {
	return &UserDatabase{sqlHandler}
}

// FindByID ...
func (db *UserDatabase) FindByID(id string) (*model.User, error) {
	userAuthTable := &userDB.UserAuthTable
	statement := db.findStatement(userAuthTable.UserID)
	valueList := []interface{}{id}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The user doesn't exist"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	user := &model.User{}
	scanList := []interface{}{&user.ID, &user.Email, &user.PasswordDigest, &user.Category}
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	userProfileTable := &userDB.ProfileTable
	statement = NewSQLBuilder().Select(userProfileTable).Match(userProfileTable.UserID).Statement()
	valueList = []interface{}{id}
	row, err = db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	exists = row.Next()
	if !exists {
		log.Info("The user info doesn't exist")
		return user, nil
	}
	userProfile := &model.Profile{}
	scanList = generateScanList(userProfile)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	user.Profile = userProfile
	return user, nil
}

// FindByEmail ...
func (db *UserDatabase) FindByEmail(email string) (*model.User, error) {
	userAuthTable := &userDB.UserAuthTable
	statement := db.findStatement(userAuthTable.Email)
	valueList := []interface{}{email}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The user doesn't exist"
		log.Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	user := &model.User{}
	scanList := []interface{}{&user.ID, &user.Email, &user.PasswordDigest, &user.Category}
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	userProfileTable := &userDB.ProfileTable
	statement = NewSQLBuilder().Select(userProfileTable).Match(userProfileTable.UserID).Statement()
	valueList = []interface{}{user.ID}
	row, err = db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer row.Close()
	exists = row.Next()
	if !exists {
		log.Error("The user info doesn't exist")
		return user, nil
	}
	userProfile := &model.Profile{}
	scanList = generateScanList(userProfile)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	user.Profile = userProfile
	return user, nil
}

func (db *UserDatabase) FindByAccountID(accountID string) (*model.Profile, error) {
	userProfileTable := &userDB.ProfileTable
	statement := NewSQLBuilder().Select(userProfileTable).Match(userProfileTable.AccountID).Statement()
	row, err := db.Query(statement, accountID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	exists := row.Next()
	if !exists {
		log.Info("The user info doesn't exist")
		return nil, nil
	}
	profile := &model.Profile{}
	scanList := generateScanList(profile)
	err = row.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return profile, nil
}

func (db *UserDatabase) FindCardByUserID(userID string) (*model.Card, error) {
	cardTable := &userDB.CardTable
	statement := NewSQLBuilder().Select(cardTable).Match(cardTable.UserID).Statement()
	rows, err := db.Query(statement, userID)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	exists := rows.Next()
	if !exists {
		err = errs.NotFound.New("Card is not found")
		return nil, err
	}
	card := &model.Card{}
	scanList := generateScanList(card)
	err = rows.Scan(scanList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	return card, nil
}

// FindListByLikeBrandID ...
func (db *UserDatabase) FindListByLikeBrandID(id string) ([]model.User, error) {
	userProfileTable := &userDB.ProfileTable
	brandLikeTable := &brandDB.BrandLikeTable
	statement := NewSQLBuilder().Select(customeUserTable).
		LeftJoinWithColumns(userProfileTable, userProfileTable.UserID, customeUserTable.UserID).
		RightJoin(brandLikeTable, brandLikeTable.UserID, customeUserTable.UserID).
		Match(brandLikeTable.BrandID).
		Statement()
	valueList := []interface{}{id}
	rows, err := db.Query(statement, valueList...)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	userList := []model.User{}
	for rows.Next() {
		user := model.User{}
		scanList := generateScanList(&user)
		err := rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		userList = append(userList, user)
	}
	return userList, nil
}

// Save ...
func (db *UserDatabase) Save(user *model.User) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Save transaction partner
		transactionPartnerTable := &transactionDB.TransactionPartnersTable
		statement := NewSQLBuilder().Insert(transactionPartnerTable)
		valueList := []interface{}{user.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Create parent user
		usersTable := &userDB.UsersTable
		statement = NewSQLBuilder().Insert(usersTable)
		valueList = []interface{}{user.ID}
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Save user cateogry
		tableQuery, columnList := db.getCategoryTableQueryAndColumnList(user.Category)
		columnQuery := generateInsertColumnQuery(columnList)
		valuesQuery := generateValuesQuery(len(columnList))
		statement = strings.Join([]string{
			rw.InsertInto, tableQuery, columnQuery,
			rw.Values, valuesQuery,
		}, " ")
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Save user auth
		userAuthTable := &userDB.UserAuthTable
		statement = NewSQLBuilder().Insert(userAuthTable)
		valueList = append(valueList, user.Email, user.PasswordDigest)
		_, err = db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// SaveProfile ...
func (db *UserDatabase) SaveProfile(user *model.User) error {
	_, err := db.Transaction(func() (interface{}, error) {
		userProfileTable := &userDB.ProfileTable
		statement := NewSQLBuilder().Insert(userProfileTable)
		valueList := generateValueList(user.Profile)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Update image to storage
		storage.UploadFileForCreate(user.Profile.Logo, storage.IconPath)
		// Update image to storage
		storage.UploadFileForCreate(user.Profile.Icon, storage.BackGroundPath)
		return nil, nil
	})
	return err
}

func (db *UserDatabase) SaveCard(card *model.Card) error {
	_, err := db.Transaction(func() (interface{}, error) {
		cardTable := &userDB.CardTable
		statement := NewSQLBuilder().Insert(cardTable)
		valueList := generateValueList(card)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

func (db *UserDatabase) DeleteCard(card *model.Card) error {
	_, err := db.Transaction(func() (interface{}, error) {
		cardTable := &userDB.CardTable
		statement := NewSQLBuilder().Delete(cardTable).Match(cardTable.ID).Statement()
		_, err := db.Exec(statement, card.ID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// UpdatePassword ...
func (db *UserDatabase) UpdatePassword(user *model.User) error {
	_, err := db.Transaction(func() (interface{}, error) {
		userAuthTable := &userDB.UserAuthTable
		statement := strings.Join([]string{
			rw.Update, userAuthTable.NAME(),
			rw.Set, userAuthTable.PasswordDigest, "= ?",
			rw.Where, userAuthTable.UserID, "= ?",
		}, " ")
		_, err := db.Exec(statement, user.PasswordDigest, user.ID)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// UpdateToActive ...
func (db *UserDatabase) UpdateToActive(user *model.User) error {
	_, err := db.Transaction(func() (interface{}, error) {
		// Save user active
		usersActiveTable := &userDB.UsersActiveTable
		statement := NewSQLBuilder().Insert(usersActiveTable)
		valueList := []interface{}{user.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
		}
		return nil, nil
	})
	return err
}

// UpdateProfile ...
func (db *UserDatabase) UpdateProfile(user *model.User) error {
	_, err := db.Transaction(func() (interface{}, error) {
		userProfileTable := &userDB.ProfileTable
		statement := NewSQLBuilder().Update(userProfileTable)
		valueList := generateValueList(user.Profile)
		valueList = append(valueList, user.ID)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		// Update image to storage
		storage.UploadFileForUpdate(user.Profile.Logo, storage.IconPath)
		// Update image to storage
		storage.UploadFileForUpdate(user.Profile.Icon, storage.BackGroundPath)
		return nil, nil
	})
	return err
}

// getCategoryTableQueryAndColumnList ...
func (db *UserDatabase) getCategoryTableQueryAndColumnList(category int) (tableQuery string, columnList []string) {
	switch category {
	case brandOwnerID:
		brandOwnersTable := &userDB.BrandOwnersTable
		tableQuery = brandOwnersTable.NAME()
		columnList = []string{brandOwnersTable.UserID}
	case partnerID:
		partnersTable := &userDB.PartnersTable
		tableQuery = partnersTable.NAME()
		columnList = []string{partnersTable.UserID}
	default:
	}
	return
}

func (db *UserDatabase) findStatement(condition string) string {
	usersActiveTable := &userDB.UsersActiveTable
	userAuthTable := &userDB.UserAuthTable
	brandOwnersTable := &userDB.BrandOwnersTable
	partnersTable := &userDB.PartnersTable
	columnList := generateColumnList(
		userAuthTable.UserID,
		userAuthTable.Email,
		userAuthTable.PasswordDigest,
		"category.category",
	)
	columnQuery := generateSelectColumnQuery(columnList)
	statement := strings.Join([]string{
		rw.Select, columnQuery,
		rw.From, userAuthTable.NAME(),
		rw.LeftJoin, usersActiveTable.NAME(),
		rw.On, userAuthTable.UserID, "=", usersActiveTable.UserID,
		rw.LeftJoin, "(",
		rw.Select,
		brandOwnersTable.UserID, ",",
		"1", rw.As, "category",
		rw.From, brandOwnersTable.NAME(),
		rw.UnionAll,
		rw.Select,
		partnersTable.UserID, ",",
		"2", rw.As, "category",
		rw.From, partnersTable.NAME(),
		")", rw.As, "category",
		rw.On, userAuthTable.UserID, "=", "category.user_id",
		rw.Where, condition, "= ?",
	}, " ")
	return statement
}
