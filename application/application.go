package application

import "github.com/usagiga/Distable/entity"

type EmojiSyncApplication interface {
	Sync(tgtServs []entity.ServerContext) (err error)
}