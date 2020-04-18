package domain

import "github.com/usagiga/distable/entity"

// EmojiArrayModel is model treating `[]entity.EmojiContext`,
// used to sync server's emoji.
type EmojiArrayModel interface {
	// Unique gets all `src` elements not included in `dst`.
	Unique(dst []entity.EmojiContext, src []entity.EmojiContext) (uniqueDst []entity.EmojiContext)
	// Unite gets all emojis without duplicate in specified emojis.
	Unite(serverEmojisArr [][]entity.EmojiContext) (unitedEmojis []entity.EmojiContext)
}

// ServerArrayModel is model treating `[]entity.ServerContext`,
// used to recognize server's type (master or slave).
type ServerArrayModel interface {
	// GetMasters gets all master servers.
	GetMasters(servers []entity.ServerContext) (masters []entity.ServerContext)
	// GetSlaves gets all slave servers.
	GetSlaves(servers []entity.ServerContext) (slaves []entity.ServerContext)
}
