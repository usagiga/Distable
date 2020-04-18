package application

import (
	"github.com/usagiga/distable/domain"
	"github.com/usagiga/distable/entity"
	"github.com/usagiga/distable/infrastructure"
	"log"
)

// EmojiSyncApplicationImpl is struct implemented `EmojiSyncApplication`.
type EmojiSyncApplicationImpl struct {
	emojiArrayModel domain.EmojiArrayModel
	serverArrayModel domain.ServerArrayModel
	emojiInfra      infrastructure.EmojiInfra
}

// NewEmojiSyncApplication initializes `EmojiSyncApplication`.
func NewEmojiSyncApplication(
	emojiArrayModel domain.EmojiArrayModel,
	serverArrayModel domain.ServerArrayModel,
	emojiInfra infrastructure.EmojiInfra,
) EmojiSyncApplication {
	return &EmojiSyncApplicationImpl{
		emojiArrayModel: emojiArrayModel,
		serverArrayModel: serverArrayModel,
		emojiInfra:      emojiInfra,
	}
}

// Sync syncs emoji among specified servers.
func (e *EmojiSyncApplicationImpl) Sync(tgtServs []entity.ServerContext) (err error) {
	masters := e.serverArrayModel.GetMasters(tgtServs)
	slaves := e.serverArrayModel.GetSlaves(tgtServs)

	// Load all server inventories
	masterInvs, err := e.getInventories(masters)
	if err != nil {
		return err
	}

	slaveInvs, err := e.getInventories(slaves)
	if err != nil {
		return err
	}

	// Get master's emojis
	var masterEmojisArr [][]entity.EmojiContext
	for _, inv := range masterInvs {
		masterEmojisArr = append(masterEmojisArr, inv.EmojiContexts)
	}

	// Unite master's emojis. It will be synced.
	applyingEmojis := e.emojiArrayModel.Unite(masterEmojisArr)

	// Apply to master servers
	for _, inv := range masterInvs {
		err = e.applyEmojis(&inv, applyingEmojis)
		if err != nil {
			return err
		}
	}

	// Apply to slave servers
	for _, inv := range slaveInvs {
		err = e.applyEmojis(&inv, applyingEmojis)
		if err != nil {
			return err
		}
	}

	return nil
}

func (e *EmojiSyncApplicationImpl) getInventory(tgtServ *entity.ServerContext) (inventory *entity.ServerEmojiInventory, err error) {
	emojis, err := e.emojiInfra.FetchAllContext(*tgtServ)
	if err != nil {
		return nil, err
	}

	inventory = &entity.ServerEmojiInventory{
		ServerContext: *tgtServ,
		EmojiContexts: emojis,
	}

	return inventory, nil
}

func (e *EmojiSyncApplicationImpl) getInventories(tgtServs []entity.ServerContext) (inventories []entity.ServerEmojiInventory, err error) {
	inventories = []entity.ServerEmojiInventory{}

	for _, tgtServ := range tgtServs {
		inventory, err := e.getInventory(&tgtServ)
		if err != nil {
			return nil, err
		}

		inventories = append(inventories, *inventory)
	}
	return inventories, nil
}

func (e *EmojiSyncApplicationImpl) applyEmojis(tgtServInv *entity.ServerEmojiInventory, applyingEmojis []entity.EmojiContext) (err error) {
	dstID := tgtServInv.ServerContext.GuildID
	dstServCtx := tgtServInv.ServerContext
	dstEmojiCtxs := tgtServInv.EmojiContexts
	addingEmojiCtxs := e.emojiArrayModel.Unique(dstEmojiCtxs, applyingEmojis) // Does't have but included in an applying array
	deletingEmojiCtxs := e.emojiArrayModel.Unique(applyingEmojis, dstEmojiCtxs) // Already had but not included in an applying array

	// Fetch adding actual images
	// TODO : Memorize already loaded emoji to reduce discord CDN server's load
	addingEmojis, err := e.emojiInfra.FetchAll(addingEmojiCtxs)
	if err != nil {
		return err
	}

	// Add all adding emojis
	for _, emoji := range addingEmojis {
		err = e.emojiInfra.Add(&emoji, &dstServCtx)
		if err != nil {
			return err
		}

		log.Printf("Added %s into %s\n", emoji.Name, dstID)
	}

	// Delete all emojis to delete
	for _, eCtx := range deletingEmojiCtxs {
		err = e.emojiInfra.Delete(&eCtx, &dstServCtx)
		if err != nil {
			return err
		}

		log.Printf("Deleted %s from %s\n", eCtx.Name, dstID)
	}

	return nil
}
