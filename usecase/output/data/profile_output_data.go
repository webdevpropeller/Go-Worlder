package outputdata

// Profile ...
type Profile struct {
	Activity   string
	Industry   string
	Company    string
	Country    string
	Address1   string
	Address2   string
	City       string
	State      string
	ZipCode    string
	URL        string
	Phone      string
	AccountID  string
	Logo       string
	Language   string
	FirstName  string
	MiddleName string
	FamilyName string
	Icon       string
}

type Option struct {
	Value string
	Text  string
}

type ProfileSelectItem struct {
	Industries    []Option
	Countries     []Option
	CardCompanies []Option
}
