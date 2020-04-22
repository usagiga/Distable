package entity

// OnUpdatedEmojiHandler is a handler to do something on updated emojis.
type OnUpdatedEmojiHandler func(srcServ *ServerContext) (err error)
