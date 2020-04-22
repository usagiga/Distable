package application

import (
	"context"
	"github.com/usagiga/distable/entity"
)

// EmojiSyncApplication is to sync emoji among specified servers.
type EmojiSyncApplication interface {
	// Sync syncs emoji among specified servers.
	Sync(tgtServs []entity.ServerContext) (err error)
}

// EmojiSyncStreamApplication is to sync emoji
// among specified servers on every updating emojis.
type EmojiSyncStreamApplication interface {
	WatchUpdatingEmoji(ctx context.Context, tgtServs []entity.ServerContext) (err error)
}
