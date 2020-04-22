package entity

// ServerContext is included GuildID on Discord to use some API.
// Also it is DAO (Included in `entity.Config`).
type ServerContext struct {
	// GuildID is ID of guild on Discord.
	// A.k.a. Server ID.
	GuildID string `json:"guild_id"`
	// ServerType means the server is master or slave.
	// If you leave it not decided, it means master.
	ServerType ServerType `json:"type"`
}

// Equals compare between this and specific `ServerContext`.
// If it seems those are equal by its own *`GuildID`*, it returns `true`.
func (sc *ServerContext) Equals(dst *ServerContext) bool {
	return sc.GuildID == dst.GuildID
}
