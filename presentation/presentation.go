package presentation

import "github.com/usagiga/Distable/entity"

type EmojiSyncCommand interface {
	Sync(servers []entity.ServerContext)
}
