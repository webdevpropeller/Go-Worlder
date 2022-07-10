package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// SearchInputPort ...
type SearchInputPort interface {
	SearchProduct(*inputdata.Search) ([]outputdata.Product, error)
}
