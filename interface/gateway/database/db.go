package database

// DB ...
type DB interface {
	NAME() string
}

// db ...
type db struct {
	name string
}

// NAME ...
func (db db) NAME() string {
	return db.name
}
