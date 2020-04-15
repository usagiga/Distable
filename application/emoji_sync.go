package application

import (
	"fmt"
	"github.com/usagiga/Distable/domain"
	"github.com/usagiga/Distable/entity"
	"github.com/usagiga/Distable/infrastructure"
	"log"
)

// EmojiSyncApplicationImpl is struct implemented `EmojiSyncApplication`.
type EmojiSyncApplicationImpl struct {
	emojiArrayModel domain.EmojiArrayModel
	emojiInfra      infrastructure.EmojiInfra
}

// NewEmojiSyncApplication initializes `EmojiSyncApplication`.
func NewEmojiSyncApplication(
	emojiArrayModel domain.EmojiArrayModel,
	emojiInfra infrastructure.EmojiInfra,
) EmojiSyncApplication {
	return &EmojiSyncApplicationImpl{
		emojiArrayModel: emojiArrayModel,
		emojiInfra:      emojiInfra,
	}
}

// Sync syncs emoji among specified servers.
func (e *EmojiSyncApplicationImpl) Sync(tgtServs []entity.ServerContext) (err error) {
	// Load all server inventories
	var inventories []entity.ServerEmojiInventory
	for _, tgtServ := range tgtServs {
		emojis, err := e.emojiInfra.FetchAllContext(tgtServ)
		if err != nil {
			return err
		}

		inventory := entity.ServerEmojiInventory{
			ServerContext: tgtServ,
			EmojiContexts: emojis,
		}

		inventories = append(inventories, inventory)
	}

	// Sync emoji
	for _, dstInv := range inventories {
		dstServ := dstInv.ServerContext

		for _, srcInv := range inventories {
			dstID := dstInv.ServerContext.GuildID
			srcID := srcInv.ServerContext.GuildID
			dstEmojis := dstInv.EmojiContexts
			srcEmojis := srcInv.EmojiContexts

			// Skip sync itself.
			if dstInv.Equals(&srcInv) {
				continue
			}

			// Get emojis needed to sync.
			// TODO : Memorize already loaded emoji to reduce discord CDN server's load
			addingEmojiCtxs := e.emojiArrayModel.Unique(dstEmojis, srcEmojis)
			addingEmoji, err := e.emojiInfra.FetchAll(addingEmojiCtxs)
			if err != nil {
				return err
			}

			// Add emojis into a destination server.
			for _, emoji := range addingEmoji {
				// Write processing log
				logMsg := fmt.Sprintf("%s => [ %s ] => %s", srcID, emoji.Name, dstID)
				log.Println(logMsg)

				err = e.emojiInfra.Add(&emoji, &dstServ)
				if err != nil {
					return err
				}
			}

			// To avoid to sync redundantly,
			// add already processed emoji into destination inventory.
			dstInv.EmojiContexts = append(dstInv.EmojiContexts, addingEmojiCtxs...)
		}
	}

	return nil
}
