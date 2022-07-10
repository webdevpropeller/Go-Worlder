package inputport

import (
	inputdata "go_worlder_system/usecase/input/data"
	outputdata "go_worlder_system/usecase/output/data"
)

// ChatInputPort ...
type ChatInputPort interface {
	IndexDestination(string) ([]outputdata.PublicUser, error)
	Show(id, userID string) ([]outputdata.ChatMessage, error)
	Send(*inputdata.ChatMessage) (*outputdata.ChatMessage, error)
}
