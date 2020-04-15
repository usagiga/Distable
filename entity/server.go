package entity

// ServerContext is included access token or ID on Discord to use some API.
// Also it is DAO (Included in `entity.Config`).
type ServerContext struct {
	// AuthorizationToken is Discord API (for bot) token.
	// You can obtain it through OAuth using bot's Client ID.
	AuthorizationToken string `json:"auth_token"`
	// GuildID is ID of guild on Discord.
	// A.k.a. Server ID.
	GuildID string `json:"guild_id"`
}

// GetBearerToken generates the body of `Authorization` HTTP header.
func (sc *ServerContext) GetBearerToken() string {
	return "Bot " + sc.AuthorizationToken
}

// Equals compare between this and specific `ServerContext`.
// If it seems those are equal by its own *`GuildID`*, it returns `true`.
func (sc *ServerContext) Equals(dst *ServerContext) bool {
	return sc.GuildID == dst.GuildID
}
