package entity

// Config is DAO represented `config.json`
type Config struct {
	Servers []ServerContext `json:"servers"`
}