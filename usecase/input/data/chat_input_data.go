package inputdata

// ChatMessageRequest ...
type ChatMessageRequest struct {
	RoomID  string
	Content string
}

// ChatMessage ...
type ChatMessage struct {
	SenderID   string
	ReceiverID string
	Content    string
}
