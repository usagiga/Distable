package application

import (
	"fmt"
	"github.com/usagiga/Distable/domain"
	"github.com/usagiga/Distable/entity"
	"github.com/usagiga/Distable/infrastructure"
	"log"
)

type EmojiSyncApplicationImpl struct {
	emojiArrayModel domain.EmojiArrayModel
	emojiInfra      infrastructure.EmojiInfra
}

func NewEmojiSyncApplication(
	emojiArrayModel domain.EmojiArrayModel,
	emojiInfra infrastructure.EmojiInfra,
) EmojiSyncApplication {
	return &EmojiSyncApplicationImpl{
		emojiArrayModel: emojiArrayModel,
		emojiInfra: emojiInfra,
	}
}

func (e *EmojiSyncApplicationImpl) Sync(tgtServs []entity.ServerContext) (err error) {
	// Load server inventories
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
	for i, srcInv := range inventories {
		// TODO : Memorize already loaded emoji to reduce discord CDN server's load
		tgtServs := srcInv.ServerContext
		for _, dstInv := range inventories {
			srcID := srcInv.ServerContext.GuildID
			dstID := dstInv.ServerContext.GuildID
			srcEmojis := srcInv.EmojiContexts
			dstEmojis := dstInv.EmojiContexts

			// Skip sync itself
			if srcInv.Equals(&dstInv) {
				continue
			}

			// Get emojis needed to sync
			addingEmojiCtxs := e.emojiArrayModel.Unique(srcEmojis, dstEmojis)
			addingEmoji, err := e.emojiInfra.FetchAll(addingEmojiCtxs)
			if err != nil {
				return err
			}

			// Add emojis into a source server
			for _, emoji := range addingEmoji {
				// Write processing log
				logMsg := fmt.Sprintf("%s => [ %s ] => %s", dstID, emoji.Name, srcID)
				log.Println(logMsg)

				err = e.emojiInfra.Add(&emoji, &tgtServs)
				if err != nil {
					return err
				}
			}

			// To avoid to sync redundantly,
			// add already processed emoji into source inventory
			srcInv.EmojiContexts = append(srcInv.EmojiContexts, addingEmojiCtxs...)
			inventories[i].EmojiContexts = append(inventories[i].EmojiContexts, addingEmojiCtxs...)
		}
	}

	return nil
}
