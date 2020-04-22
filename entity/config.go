package entity

// Config is DAO represented `config.json`.
type Config struct {
	// Credential is used to connect to Discord API (for bot) token.
	Credential *Credential `json:"credential"`
	// Servers represents target servers to sync emojis.
	Servers []ServerContext `json:"servers"`
}
