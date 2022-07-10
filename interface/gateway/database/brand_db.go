package database

// BrandDB ...
type BrandDB struct {
	db
	BrandsTable        BrandsTable
	BrandsActiveTable  BrandsActiveTable
	BrandsDeletedTable BrandsDeletedTable
	BrandLikeTable     BrandLikeTable
}

// BrandsTable ...
type BrandsTable struct {
	table
	ID          string
	UserID      string
	CategoryID  string
	Name        string
	Slogan      string
	LogoImage   string
	Description string
	IsDraft     string
	CreatedAt   string
	UpdatedAt   string
}

// BrandsActiveTable ...
type BrandsActiveTable struct {
	table
	BrandID   string
	CreatedAt string
}

// BrandsDeletedTable ...
type BrandsDeletedTable struct {
	table
	BrandID   string
	CreatedAt string
}

// BrandLikeTable ...
type BrandLikeTable struct {
	table
	BrandID string
	UserID  string
}

// NewBrandDB ...
func NewBrandDB() *BrandDB {
	brandDB := &BrandDB{}
	initialize(brandDB)
	return brandDB
}
