package infrastructure

import (
	"context"
	"github.com/usagiga/distable/entity"
)

// EmojiInfra treats emojis through Discord API / CDN server.
type EmojiInfra interface {
	// Add adds emoji into a specific server.
	Add(destServCtx *entity.ServerContext, emoji *entity.Emoji) (err error)
	// Fetch fetches actual emoji image from its context.
	Fetch(emojiCtx *entity.EmojiContext) (emoji *entity.Emoji, err error)
	// FetchAll fetches actual emoji image from its context array one-by-one.
	FetchAll(emojiCtxs []entity.EmojiContext) (emojis []entity.Emoji, err error)
	// FetchAllContext fetches emoji contexts on a specific server.
	FetchAllContext(servCtx entity.ServerContext) (emojis []entity.EmojiContext, err error)
	// Delete delets the emoji from the server.
	Delete(destServCtx *entity.ServerContext, emojiCtx *entity.EmojiContext) (err error)
}

// EmojiStreamInfra treats the events on Discord API server.
// Particularly about updating emojis.
type EmojiStreamInfra interface {
	// WatchUpdatingEmoji waits updating emoji on specified servers
	// and notifies about it to the handler.
	WatchUpdatingEmoji(ctx context.Context, handler entity.OnUpdatedEmojiHandler) (err error)
}
