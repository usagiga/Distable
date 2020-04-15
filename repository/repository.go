package repository

import "github.com/usagiga/Distable/entity"

type ConfigRepository interface {
	Get() *entity.Config
}
