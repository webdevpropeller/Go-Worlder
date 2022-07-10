package database

// ChatDB ...
type ChatDB struct {
	db
	ChatMessagesTable ChatMessagesTable
}

// ChatMessagesTable ...
type ChatMessagesTable struct {
	table
	ID         string
	SenderID   string
	ReceiverID string
	Content    string
	CreatedAt  string
}

// NewChatDB ...
func NewChatDB() *ChatDB {
	chatDB := &ChatDB{}
	initialize(chatDB)
	return chatDB
}
