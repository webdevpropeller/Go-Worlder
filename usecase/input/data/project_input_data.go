package inputdata

// Project ...
type Project struct {
	UserID string
	Name   string
}

// NewProject ...
type NewProject struct {
	*Project
}

// UpdatedProject ...
type UpdatedProject struct {
	ID string
	*Project
}
