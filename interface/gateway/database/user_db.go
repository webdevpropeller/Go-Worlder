package database

// UserDB ...
type UserDB struct {
	db
	UsersTable       UsersTable
	BrandOwnersTable BrandOwnersTable
	PartnersTable    PartnersTable
	UsersActiveTable UsersActiveTable
	UsersLeftTable   UsersLeftTable
	UserTokensTable  UserTokensTable
	UserAuthTable    UserAuthTable
	ProfileTable     ProfileTable
	CardTable        CardTable
}

// UsersTable ...
type UsersTable struct {
	table
	ID        string
	CreatedAt string
}

// BrandOwnersTable ...
type BrandOwnersTable struct {
	table
	UserID string
}

// PartnersTable ...
type PartnersTable struct {
	table
	UserID string
}

// UsersActiveTable ...
type UsersActiveTable struct {
	table
	UserID    string
	CreatedAt string
}

// UsersLeftTable ...
type UsersLeftTable struct {
	table
	UserID    string
	CreatedAt string
}

// UserTokensTable ...
type UserTokensTable struct {
	table
	UserID    string
	Token     string
	CreatedAt string
}

// UserAuthTable ...
type UserAuthTable struct {
	table
	UserID         string
	Email          string
	PasswordDigest string
	CreatedAt      string
	UpdatedAt      string
}

// ProfileTable ...
type ProfileTable struct {
	table
	UserID     string
	ActivityID string
	IndustryID string
	Company    string
	CountryID  string
	Address1   string
	Address2   string
	ZipCode    string
	URL        string
	Phone      string
	AccountID  string
	Logo       string
	FirstName  string
	MiddleName string
	FamilyName string
	Icon       string
}

type CardTable struct {
	table
	ID           string
	UserID       string
	CompanyID    string
	Owner        string
	Number       string
	Expiry       string
	SecurityCode string
}

// NewUserDB ...
func NewUserDB() *UserDB {
	userDB := &UserDB{}
	initialize(userDB)
	return userDB
}
