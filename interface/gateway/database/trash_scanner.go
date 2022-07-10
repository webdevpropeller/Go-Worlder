package database

// TrashScanner is dummy structure for truncating unnecessary columns when scanning
type trashScanner struct{}

// Scan ...
func (trashScanner) Scan(interface{}) error {
	return nil
}
