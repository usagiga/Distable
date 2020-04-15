package entity

// Config is DAO represented `config.json`.
type Config struct {
	// Servers represents target servers to sync emojis.
	Servers []ServerContext `json:"servers"`
}
