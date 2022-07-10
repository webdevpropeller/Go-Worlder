package localcache

import (
	"go_worlder_system/domain/model"

	log "github.com/sirupsen/logrus"
)

// ChatMessageLocalCache ...
type ChatMessageLocalCache struct {
	Cache
}

// NewChatMessageLocalCache ...
func NewChatMessageLocalCache(cache Cache) *ChatMessageLocalCache {
	return &ChatMessageLocalCache{cache}
}

// FindByID ...
func (lc *ChatMessageLocalCache) FindByID(id string, f func(string) (*model.ChatMessage, error)) (*model.ChatMessage, error) {
	if msg, ok := lc.Get(id); ok {
		return msg.(*model.ChatMessage), nil
	}
	msg, err := f(id)
	if err != nil {
		log.Error(err)
		return nil, err
	}
	lc.Set(msg.ID, msg, lc.DefaultExpiration())
	return msg, err
}

// Save ...
func (lc *ChatMessageLocalCache) Save(msg *model.ChatMessage) error {
	lc.Set(msg.ID, msg, lc.DefaultExpiration())
	return nil
}
