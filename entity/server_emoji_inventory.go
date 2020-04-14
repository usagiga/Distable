package entity

type ServerEmojiInventory struct {
	ServerContext ServerContext
	EmojiContexts []EmojiContext
}

func (s *ServerEmojiInventory) Equals(dst *ServerEmojiInventory) bool {
	return s.ServerContext.Equals(&dst.ServerContext)
}