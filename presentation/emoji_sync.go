package presentation

import (
	"github.com/usagiga/Distable/application"
	"github.com/usagiga/Distable/entity"
	"log"
	"os"
)

// EmojiSyncCommandImpl is struct implemented `EmojiSyncCommand`.
type EmojiSyncCommandImpl struct {
	emojiSyncApp application.EmojiSyncApplication
}

// NewEmojiSyncCommand initializes `EmojiSyncCommand`
func NewEmojiSyncCommand(emojiSyncApp application.EmojiSyncApplication) EmojiSyncCommand {
	return &EmojiSyncCommandImpl{
		emojiSyncApp: emojiSyncApp,
	}
}

// Sync let Distable sync emoji among specified servers.
func (e *EmojiSyncCommandImpl) Sync(tgtServs []entity.ServerContext) {
	err := e.emojiSyncApp.Sync(tgtServs)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("EmojiSync.Sync(): Error raised: %s", errMsg)
		os.Exit(1)
	}
}
