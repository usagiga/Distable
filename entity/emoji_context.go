package entity

// EmojiContext is overview of emoji without its actual image.
type EmojiContext struct {
	// ID is emoji id on Discord.
	ID string
	// Name is its screen name.
	Name string
	RequireColons bool
	Animated bool
}

// Equals compare between this and specific `EmojiContext`.
// If it seems those are equal by its own *`Name`*, it returns `true`.
func (e *EmojiContext) Equals(dst *EmojiContext) bool {
	return e.Name == dst.Name
}

func (e *EmojiContext) GetExtension() string {
	if e.Animated {
		return "gif"
	}

	return "png"
}