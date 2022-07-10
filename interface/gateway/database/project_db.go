package database

// ProjectDB ...
type ProjectDB struct {
	db
	ProjectsTable        ProjectsTable
	ProjectsActiveTable  ProjectsActiveTable
	ProjectsDeletedTable ProjectsDeletedTable
	ProjectsFormalTable  ProjectsFormalTable
	ProjectsOfDraftTable ProjectsOfDraftTable
}

// ProjectsTable ...
type ProjectsTable struct {
	table
	ID        string
	UserID    string
	Name      string
	CreatedAt string
	UpdatedAt string
}

// ProjectsActiveTable ...
type ProjectsActiveTable struct {
	table
	ProjectID string
	CreatedAt string
}

// ProjectsFormalTable ...
type ProjectsFormalTable struct {
	table
	ProjectID string
	CreatedAt string
}

// ProjectsOfDraftTable ...
type ProjectsOfDraftTable struct {
	table
	ProjectID string
	CreatedAt string
}

// ProjectsDeletedTable ...
type ProjectsDeletedTable struct {
	table
	ProjectID string
	CreatedAt string
}

// NewProjectDB ...
func NewProjectDB() *ProjectDB {
	projectDB := &ProjectDB{}
	initialize(projectDB)
	return projectDB
}
