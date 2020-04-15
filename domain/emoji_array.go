package domain

import (
	"github.com/usagiga/Distable/entity"
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
