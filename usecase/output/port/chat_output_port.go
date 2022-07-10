//go:generate mockgen -source=$GOFILE -destination=../mock_port/mock_$GOFILE
package outputport

import (
	"go_worlder_system/domain/model"
	outputdata "go_worlder_system/usecase/output/data"
)

// ChatOutputPort ...
type ChatOutputPort interface {
	IndexDestination([]model.Profile) []outputdata.PublicUser
	Show([]model.ChatMessage) []outputdata.ChatMessage
	Send(*model.ChatMessage) *outputdata.ChatMessage
}
