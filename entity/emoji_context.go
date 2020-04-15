package entity

// EmojiContext is overview of emoji without its actual image.
type EmojiContext struct {
	// ID is emoji id on Discord.
	ID string
	// Name is its screen name.
	Name string
	// RequireColons is decided behavior on chat in Discord.
	// It is not applied synced server currently.
	RequireColons bool
	// Animated is represents the emoji is animated emoji(gif) or not.
	Animated bool
}

// Equals compare between this and specific `EmojiContext`.
// If it seems those are equal by its own *`Name`*, it returns `true`.
func (e *EmojiContext) Equals(dst *EmojiContext) bool {
	return e.Name == dst.Name
}

// GetExtension returns the emoji's image file extension.
// Its value is only `png` or `gif`.
func (e *EmojiContext) GetExtension() string {
	if e.Animated {
		return "gif"
	}

	return "png"
}
