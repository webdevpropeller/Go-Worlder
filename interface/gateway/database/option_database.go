package database

import (
	"go_worlder_system/domain/model"

	log "github.com/sirupsen/logrus"
)

type OptionDatabase struct {
	SQLHandler
}

func NewOptionDatabase(sqlHandler SQLHandler) *OptionDatabase {
	return &OptionDatabase{sqlHandler}
}

func (db *OptionDatabase) FindIndustries() (model.Industries, error) {
	mIndustriesTable := masterDB.MIndustriesTable
	statement := NewSQLBuilder().Select(mIndustriesTable).Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	industries := model.Industries{}
	for rows.Next() {
		industry := model.Option{}
		scanList := generateScanList(&industry)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		industries = append(industries, industry)
	}
	return industries, nil
}

func (db *OptionDatabase) FindCountries() (model.Countries, error) {
	mCountriesTable := masterDB.MCountriesTable
	statement := NewSQLBuilder().Select(mCountriesTable).Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	countries := model.Countries{}
	for rows.Next() {
		country := model.Option{}
		scanList := generateScanList(&country)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		countries = append(countries, country)
	}
	return countries, nil
}

func (db *OptionDatabase) FindCardCompanies() (model.CardCompanies, error) {
	mCardCompaniesTable := masterDB.MCardCompaniesTable
	statement := NewSQLBuilder().Select(mCardCompaniesTable).Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	cardCompanies := model.CardCompanies{}
	for rows.Next() {
		cardCompany := model.Option{}
		scanList := generateScanList(&cardCompany)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		cardCompanies = append(cardCompanies, cardCompany)
	}
	return cardCompanies, nil
}

func (db *OptionDatabase) FindCategories() (model.Categories, error) {
	mCategoriesTable := masterDB.MCategoriesTable
	statement := NewSQLBuilder().Select(mCategoriesTable).Statement()
	rows, err := db.Query(statement)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	defer rows.Close()
	categories := model.Categories{}
	for rows.Next() {
		category := model.Option{}
		scanList := generateScanList(&category)
		err = rows.Scan(scanList...)
		if err != nil {
			log.Error(err)
			return nil, err
		}
		categories = append(categories, category)
	}
	return categories, nil
}
