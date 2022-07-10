package sqlhandler

import (
	"database/sql"
	"go_worlder_system/consts"
	"go_worlder_system/errs"
	"go_worlder_system/interface/gateway/database"
	"os"
	"strings"

	log "github.com/sirupsen/logrus"
)

// SQLHandler has a connection with DB
type SQLHandler struct {
	DB *sql.DB
}

// NewSQLHandler ...
func NewSQLHandler() *SQLHandler {
	dbms := "mysql"
	user := os.Getenv(consts.MYSQL_USER)
	password := os.Getenv(consts.MYSQL_PASSWORD)
	host := os.Getenv(consts.MYSQL_HOST)
	option := "?parseTime=true&loc=Asia%2FTokyo"
	connect := strings.Join([]string{user, ":", password, "@", "tcp(", host, ":3306)", "/user", option}, "")
	db, err := sql.Open(dbms, connect)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		panic(err)
	}
	db.SetMaxIdleConns(300)
	db.SetMaxOpenConns(300)
	return &SQLHandler{DB: db}
}

// Exec executes the SQL that manipulates the data of the table
func (handler *SQLHandler) Exec(statement string, args ...interface{}) (database.Result, error) {
	stmt, err := handler.DB.Prepare(statement)
	if err != nil {
		log.WithFields(log.Fields{
			"statement": statement,
			"args":      args,
		}).Error(err)
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	res, err := stmt.Exec(args...)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	return &SQLResult{Result: res}, nil
}

// Query gets data from the database
func (handler *SQLHandler) Query(statement string, args ...interface{}) (database.Row, error) {
	stmt, err := handler.DB.Prepare(statement)
	if err != nil {
		log.WithFields(log.Fields{}).Error(err)
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	rows, err := stmt.Query(args...)
	if err != nil {
		log.WithFields(log.Fields{
			"statement": statement,
			"args":      args,
		}).Error(err)
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	return &SQLRow{Rows: rows}, nil
}

// Transaction ...
func (handler *SQLHandler) Transaction(f func() (interface{}, error)) (interface{}, error) {
	tx, err := handler.DB.Begin()
	if err != nil {
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	v, ferr := f()
	if err != nil {
		tx.Rollback()
		return nil, ferr
	}
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		return nil, errs.Failed.Wrap(err, err.Error())
	}
	return v, nil
}

// SQLResult ...
type SQLResult struct {
	Result sql.Result
}

// LastInsertId ...
func (r *SQLResult) LastInsertId() (int64, error) {
	res, err := r.Result.LastInsertId()
	if err != nil {
		return res, errs.Failed.Wrap(err, err.Error())
	}
	return res, nil
}

// RowsAffected ...
func (r *SQLResult) RowsAffected() (int64, error) {
	res, err := r.Result.LastInsertId()
	if err != nil {
		return res, errs.Failed.Wrap(err, err.Error())
	}
	return res, nil
}

// SQLRow ...
type SQLRow struct {
	Rows *sql.Rows
}

// Scan ...
func (r *SQLRow) Scan(dest ...interface{}) error {
	if err := r.Rows.Scan(dest...); err != nil {
		return errs.Failed.Wrap(err, err.Error())
	}
	return nil
}

// Next ...
func (r *SQLRow) Next() bool {
	return r.Rows.Next()
}

// Close ...
func (r SQLRow) Close() error {
	if err := r.Rows.Close(); err != nil {
		return errs.Failed.Wrap(err, err.Error())
	}
	return nil
}
