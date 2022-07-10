package database

// NoSQLHandler ...
type NoSQLHandler interface {
	HSet(key string, field string, value interface{}) error
	HGet(key string, field string) (string, error)
}
