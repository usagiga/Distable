package entity

// Credential is used to connect to Discord API server.
type Credential struct {
	// BotToken is a Discord Bot access token.
	BotToken string `json:"bot_token"`
}

// GetBearerToken generates the body of `Authorization` HTTP header.
func (at *Credential) GetBearerToken() string {
	return "Bot " + at.BotToken
}
