package database

// SQLHandler ...
type SQLHandler interface {
	Exec(string, ...interface{}) (Result, error)
	Query(string, ...interface{}) (Row, error)
	Transaction(func() (interface{}, error)) (interface{}, error)
}

// Result ...
type Result interface {
	LastInsertId() (int64, error)
	RowsAffected() (int64, error)
}

// Row ...
type Row interface {
	Scan(...interface{}) error
	Next() bool
	Close() error
}
