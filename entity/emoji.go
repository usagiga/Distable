package entity

import "github.com/vincent-petithory/dataurl"

// Emoji is its overview and its actual image.
type Emoji struct {
	EmojiContext

	DataURI *dataurl.DataURL
}

// ToURIString returns DataURI string according to its `DataURI`
func (e *Emoji) ToURIString() string {
	return e.DataURI.String()
}
