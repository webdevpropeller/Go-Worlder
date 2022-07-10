//go:generate mockgen -source=$GOFILE -destination=../mock_repository/mock_$GOFILE
package repository

import "go_worlder_system/domain/model"

// ChatMessageRepository ...
type ChatMessageRepository interface {
	FindDestinationListByUserID(string) ([]model.Profile, error)
	FindListByUserIDs(user1ID, user2ID string) ([]model.ChatMessage, error)
	Save(*model.ChatMessage) error
}
