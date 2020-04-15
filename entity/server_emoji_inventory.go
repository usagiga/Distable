package entity

// ServerEmojiInventory represents what emojis on specific server.
type ServerEmojiInventory struct {
	ServerContext ServerContext
	EmojiContexts []EmojiContext
}

// Equals compare between this and specific `ServerEmojiInventory`.
// If it seems those are equal by its own *`ServerContext`*, it returns `true`.
func (s *ServerEmojiInventory) Equals(dst *ServerEmojiInventory) bool {
	return s.ServerContext.Equals(&dst.ServerContext)
}
