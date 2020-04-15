package presentation

import "github.com/usagiga/Distable/entity"

// EmojiSyncCommand is a view on CLI.
// It shows results of `EmojiSyncApplications`.
type EmojiSyncCommand interface {
	// Sync let Distable sync emoji among specified servers.
	Sync(servers []entity.ServerContext)
}
