package database

// CategoryDB ...
type CategoryDB struct {
	db
	MCategoriesTable MCategoriesTable
}

// MCategoriesTable ...
type MCategoriesTable struct {
	table
	ID   string
	Name string
}

// NewCategoryDB ...
func NewCategoryDB() *CategoryDB {
	categoryDB := &CategoryDB{}
	initialize(categoryDB)
	return categoryDB
}
