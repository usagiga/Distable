package application

import (
	"github.com/usagiga/Distable/domain"
	"github.com/usagiga/Distable/entity"
	"github.com/usagiga/Distable/infrastructure"
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
	for _, srcInv := range inventories {
		// TODO : Memorize already loaded emoji to reduce discord CDN server's load
		tgtServs := srcInv.ServerContext
		for _, dstInv := range inventories {
			if srcInv.Equals(&dstInv) {
				continue
			}

			addingEmojiCtxs := e.emojiArrayModel.Unique(srcInv.EmojiContexts, dstInv.EmojiContexts)
			addingEmoji, err := e.emojiInfra.FetchAll(addingEmojiCtxs)
			if err != nil {
				return err
			}

			for _, emoji := range addingEmoji {
				err = e.emojiInfra.Add(&emoji, &tgtServs)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
