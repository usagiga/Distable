package application

import "github.com/usagiga/distable/entity"

// EmojiSyncApplication is to sync emoji among specified servers.
type EmojiSyncApplication interface {
	// Sync syncs emoji among specified servers.
	Sync(tgtServs []entity.ServerContext) (err error)
}
