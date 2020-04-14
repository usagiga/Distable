package domain

import (
	"github.com/usagiga/Distable/entity"
)

type EmojiArrayModelImpl struct {}

func NewEmojiArrayModel() EmojiArrayModel {
	return &EmojiArrayModelImpl{}
}

func (e *EmojiArrayModelImpl) Unique(src []entity.EmojiContext, dst []entity.EmojiContext) (uniqueSrc []entity.EmojiContext) {
	uniqueSrc = []entity.EmojiContext{}

	for _, dstElem := range dst {
		// Check elem is included in src or not
		found := false
		for _, srcElem := range src {
			if srcElem.Equals(&dstElem) {
				found = true
				break
			}
		}

		if found {
			continue
		}

		// If it isn't included, Append it
		uniqueSrc = append(uniqueSrc, dstElem)
	}

	return uniqueSrc
}
