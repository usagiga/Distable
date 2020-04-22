package application

import (
	"context"
	"github.com/usagiga/distable/entity"
	"github.com/usagiga/distable/infrastructure"
	"sync"
)

// EmojiSyncStreamApplicationImpl is struct implemented `EmojiSyncStreamApplication`.
type EmojiSyncStreamApplicationImpl struct {
	sync.Mutex
	emojiSyncApp     EmojiSyncApplication
	emojiStreamInfra infrastructure.EmojiStreamInfra
}

// NewEmojiSyncStreamApplication initializes `EmojiSyncStreamApplication`.
func NewEmojiSyncStreamApplication(
	emojiSyncApp EmojiSyncApplication,
	emojiStreamInfra infrastructure.EmojiStreamInfra,
) EmojiSyncStreamApplication {
	return &EmojiSyncStreamApplicationImpl{
		emojiSyncApp:     emojiSyncApp,
		emojiStreamInfra: emojiStreamInfra,
	}
}

// EmojiSyncStreamApplication is to sync emoji
// among specified servers on every updating emojis.
func (e *EmojiSyncStreamApplicationImpl) WatchUpdatingEmoji(ctx context.Context, tgtServs []entity.ServerContext) (err error) {
	onUpdatedEmoji := func(srcServ *entity.ServerContext) (err error) {
		// Mutex
		// TODO : Reduce running redundantly on updating by itself
		e.Lock()
		defer e.Unlock()

		// Apply Master & Slave to sync correctly
		for i, v := range tgtServs {
			if !v.Equals(srcServ) {
				tgtServs[i].ServerType = entity.Slave
				continue
			}

			tgtServs[i].ServerType = entity.Master
		}

		// Sync
		err = e.emojiSyncApp.Sync(tgtServs)
		if err != nil {
			return err
		}

		return nil
	}

	// Register the handler to watcher.
	err = e.emojiStreamInfra.WatchUpdatingEmoji(ctx, onUpdatedEmoji)
	if err != nil {
		return err
	}

	return nil
}
