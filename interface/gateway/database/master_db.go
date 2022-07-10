package database

type MasterDB struct {
	db
	MIndustriesTable    MIndustriesTable
	MCountriesTable     MCountriesTable
	MCardCompaniesTable MCardCompaniesTable
	MCategoriesTable    MCategoriesTable
}

type MIndustriesTable struct {
	table
	ID   string
	Name string
}

type MCountriesTable struct {
	table
	ID   string
	Name string
}

type MCardCompaniesTable struct {
	table
	ID   string
	Name string
}

type MLanguagesTable struct {
	table
	ID   string
	Name string
}

func NewMasterDB() *MasterDB {
	masterDB := &MasterDB{}
	initialize(masterDB)
	return masterDB
}
