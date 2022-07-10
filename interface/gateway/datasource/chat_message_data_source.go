package datasource

import (
	"go_worlder_system/domain/model"
	"go_worlder_system/interface/gateway/database"
	"go_worlder_system/interface/gateway/localcache"
	"go_worlder_system/interface/gateway/remotecache"
)

// ChatMessageDataSource ...
type ChatMessageDataSource struct {
	database    *database.ChatMessageDatabase
	localcache  *localcache.ChatMessageLocalCache
	remotecache *remotecache.ChatMessageRemoteCache
}

// NewChatMessageDataSource ...
func NewChatMessageDataSource(
	database *database.ChatMessageDatabase,
	localcache *localcache.ChatMessageLocalCache,
	remotecache *remotecache.ChatMessageRemoteCache,
) *ChatMessageDataSource {
	return &ChatMessageDataSource{
		database:    database,
		localcache:  localcache,
		remotecache: remotecache,
	}
}

func (ds *ChatMessageDataSource) FindDestinationListByUserID(userID string) ([]model.Profile, error) {
	return ds.database.FindDestinationListByUserID(userID)
}

// FindListByRoomID ...
func (ds *ChatMessageDataSource) FindListByUserIDs(user1ID, user2ID string) ([]model.ChatMessage, error) {
	return ds.database.FindListByUserIDs(user1ID, user2ID)
}

// Save ...
func (ds *ChatMessageDataSource) Save(msg *model.ChatMessage) error {
	return ds.database.Save(msg)
}
