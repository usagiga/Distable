package domain

import "github.com/usagiga/distable/entity"

// EmojiArrayModel is model treating `[]entity.EmojiContext`,
// used to sync server's emoji.
type EmojiArrayModel interface {
	// Unique gets all `src` elements not included in `dst`.
	Unique(dst []entity.EmojiContext, src []entity.EmojiContext) (uniqueDst []entity.EmojiContext)
}
