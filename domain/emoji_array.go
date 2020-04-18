package domain

import (
	"github.com/usagiga/distable/entity"
)

// EmojiArrayModelImpl is struct implemented `EmojiArrayModel`.
type EmojiArrayModelImpl struct{}

// NewEmojiArrayModel initializes `EmojiArrayModel`.
func NewEmojiArrayModel() EmojiArrayModel {
	return &EmojiArrayModelImpl{}
}

// Unique gets all `src` elements not included in `dst`.
func (e *EmojiArrayModelImpl) Unique(dst []entity.EmojiContext, src []entity.EmojiContext) (uniqueDst []entity.EmojiContext) {
	uniqueDst = []entity.EmojiContext{}

	// for all `src` elements,
	for _, srcElem := range src {
		// Check elem is included in dst or not.
		found := false
		for _, dstElem := range dst {
			if dstElem.Equals(&srcElem) {
				found = true
				break
			}
		}

		// If it is included, pass through.
		if found {
			continue
		}

		// If it isn't included, Append it.
		uniqueDst = append(uniqueDst, srcElem)
	}

	return uniqueDst
}

// Unite gets all emojis without duplicate in specified emojis.
func (e *EmojiArrayModelImpl) Unite(serverEmojisArr [][]entity.EmojiContext) (unitedEmojis []entity.EmojiContext) {
	unitedEmojis = []entity.EmojiContext{}

	// Unite
	for _, servEmojis := range serverEmojisArr {
		for _, e := range servEmojis {
			found := false

			// Avoid to add redundantly
			for _, ue := range unitedEmojis {
				if e.Equals(&ue) {
					found = true
					break
				}
			}

			if found {
				continue
			}

			// Append
			unitedEmojis = append(unitedEmojis, e)
		}
	}

	return unitedEmojis
}
