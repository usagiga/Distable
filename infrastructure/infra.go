package infrastructure

import "github.com/usagiga/distable/entity"

// EmojiInfra treats emojis through Discord API / CDN server.
type EmojiInfra interface {
	// Add adds emoji into a specific server.
	Add(emoji *entity.Emoji, destServCtx *entity.ServerContext) (err error)
	// Fetch fetches actual emoji image from its context.
	Fetch(emojiCtx *entity.EmojiContext) (emoji *entity.Emoji, err error)
	// FetchAll fetches actual emoji image from its context array one-by-one.
	FetchAll(emojiCtxs []entity.EmojiContext) (emojis []entity.Emoji, err error)
	// FetchAllContext fetches emoji contexts on a specific server.
	FetchAllContext(servCtx entity.ServerContext) (emojis []entity.EmojiContext, err error)
}
