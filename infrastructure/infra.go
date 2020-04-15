package infrastructure

import "github.com/usagiga/Distable/entity"

type EmojiInfra interface {
	Add(emoji *entity.Emoji, destServCtx *entity.ServerContext) (err error)
	Fetch(emojiCtx *entity.EmojiContext) (emoji *entity.Emoji, err error)
	FetchAll(emojiCtxs []entity.EmojiContext) (emojis []entity.Emoji, err error)
	FetchAllContext(servCtx entity.ServerContext) (emojis []entity.EmojiContext, err error)
}
