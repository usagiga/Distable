package domain

import "github.com/usagiga/Distable/entity"

// EmojiArrayModel is model treating `[]entity.EmojiContext`,
// used to sync server's emoji
type EmojiArrayModel interface {
	// Unique gets all `dst` elements not included in `src`
	Unique(src []entity.EmojiContext, dst []entity.EmojiContext) (uniqueSrc []entity.EmojiContext)
}