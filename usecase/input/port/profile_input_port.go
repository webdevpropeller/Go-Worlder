package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// ProfileInputPort ...
type ProfileInputPort interface {
	Show(string) (*outputdata.PublicUser, error)
	New() (*outputdata.ProfileSelectItem, error)
	Create(*inputdata.Profile) (*outputdata.User, error)
	Edit(string) (*outputdata.Profile, error)
	Update(*inputdata.Profile) error
	IndexBrandLike(string) ([]outputdata.Brand, error)
}
