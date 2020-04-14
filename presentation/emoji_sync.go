package presentation

import (
	"github.com/usagiga/Distable/application"
	"github.com/usagiga/Distable/entity"
	"log"
	"os"
)

type EmojiSyncCommandImpl struct {
	emojiSyncApp application.EmojiSyncApplication
}

func NewEmojiSyncCommand(emojiSyncApp application.EmojiSyncApplication) EmojiSyncCommand {
	return &EmojiSyncCommandImpl{
		emojiSyncApp: emojiSyncApp,
	}
}

func (e *EmojiSyncCommandImpl) Sync(tgtServs []entity.ServerContext) {
	err := e.emojiSyncApp.Sync(tgtServs)
	if err != nil {
		errMsg := err.Error()
		log.Fatalf("EmojiSync(): Error raised: %s", errMsg)
		os.Exit(1)
	}
}
