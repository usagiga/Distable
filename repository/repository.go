package repository

import "github.com/usagiga/distable/entity"

// ConfigRepository treats Distable configs.
type ConfigRepository interface {
	// Get gets Distable config from cache / file.
	// Its behavior is as singleton.
	Get() *entity.Config
}
