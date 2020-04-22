package presentation

import "github.com/usagiga/distable/entity"

// EmojiSyncCommand is a view on CLI.
// It shows results of `EmojiSyncApplications`.
type EmojiSyncCommand interface {
	// Sync lets Distable sync emoji among specified servers now.
	Sync(servers []entity.ServerContext)
}

// EmojiSyncStreamCommand is a view on CLI.
// It shows results of `EmojiSyncStreamApplications`.
type EmojiSyncStreamCommand interface {
	// RunSyncService lets Distable wait updating emojis to sync them among specified servers.
	RunSyncService(servers []entity.ServerContext)
}
