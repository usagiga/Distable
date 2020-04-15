package domain

import (
	"github.com/usagiga/Distable/entity"
)

type EmojiArrayModelImpl struct{}

func NewEmojiArrayModel() EmojiArrayModel {
	return &EmojiArrayModelImpl{}
}

func (e *EmojiArrayModelImpl) Unique(dst []entity.EmojiContext, src []entity.EmojiContext) (uniqueDst []entity.EmojiContext) {
	uniqueDst = []entity.EmojiContext{}

	for _, srcElem := range src {
		// Check elem is included in dst or not
		found := false
		for _, abcElem := range dst {
			if abcElem.Equals(&srcElem) {
				found = true
				break
			}
		}

		if found {
			continue
		}

		// If it isn't included, Append it
		uniqueDst = append(uniqueDst, srcElem)
	}

	return uniqueDst
}
