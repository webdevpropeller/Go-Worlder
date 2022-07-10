package repository

// BaseRepository ...
type BaseRepository interface {
	FindListByUserID(string) ([]interface{}, error)
	FindByID(string) (interface{}, error)
	FindUserIDByID(string) string
	Save(interface{}) error
	Update(interface{}) error
	Delete(string) error
}
