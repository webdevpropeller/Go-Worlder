package database

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/errs"
	"go_worlder_system/usecase/repository"

	log "github.com/sirupsen/logrus"
)

// UserTokenDatabase ...
type UserTokenDatabase struct {
	SQLHandler
}

// NewUserTokenDatabase ...
func NewUserTokenDatabase(sqlHandler SQLHandler) repository.UserTokenRepository {
	return &UserTokenDatabase{sqlHandler}
}

// FindByToken ...
func (db *UserTokenDatabase) FindByToken(token string) (*model.UserToken, error) {
	userTokensTable := &userDB.UserTokensTable
	statement := NewSQLBuilder().Select(userTokensTable).Match(userTokensTable.Token).Statement()
	valueList := []interface{}{token}
	row, err := db.Query(statement, valueList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	defer row.Close()
	exists := row.Next()
	if !exists {
		errMsg := "The token doesn't exist"
		log.WithFields(log.Fields{}).Error(errMsg)
		return nil, errs.NotFound.New(errMsg)
	}
	userToken := &model.UserToken{}
	user := &model.User{}
	scanList := []interface{}{&user.ID, &userToken.Token}
	err = row.Scan(scanList...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, err
	}
	userToken.User = user
	return userToken, nil
}

// Save ...
func (db *UserTokenDatabase) Save(userToken *model.UserToken) error {
	_, err := db.Transaction(func() (interface{}, error) {
		userTokensTable := &userDB.UserTokensTable
		statement := NewSQLBuilder().Insert(userTokensTable)
		valueList := generateValueList(userToken)
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}

// Delete ...
func (db *UserTokenDatabase) Delete(userToken *model.UserToken) error {
	_, err := db.Transaction(func() (interface{}, error) {
		userTokensTable := &userDB.UserTokensTable
		statement := NewSQLBuilder().Delete(userTokensTable).Match(userTokensTable.UserID).Statement()
		valueList := []interface{}{userToken.User.ID}
		_, err := db.Exec(statement, valueList...)
		if err != nil {
			log.WithFields(log.Fields{}).Error(err)
			return nil, err
		}
		return nil, nil
	})
	return err
}
