package remotecache

import "go_worlder_system/domain/model"

// ChatMessageRemoteCache ...
type ChatMessageRemoteCache struct {
}

// NewChatMessageRemoteCache ...
func NewChatMessageRemoteCache() *ChatMessageRemoteCache {
	return &ChatMessageRemoteCache{}
}

// GetByID ...
func (rc *ChatMessageRemoteCache) GetByID(id string, f func(string) (*model.ChatMessage, error)) (*model.ChatMessage, error) {
	return nil, nil
}
