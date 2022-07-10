package inputport

import inputdata "go_worlder_system/usecase/input/data"

// OrderInputPort ...
type OrderInputPort interface {
	Create(*inputdata.Order) error
}
