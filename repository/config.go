package repository

import (
	"encoding/json"
	"github.com/usagiga/Distable/entity"
	"io/ioutil"
)

type ConfigRepositoryImpl struct {
	configPath string
	configInstance *entity.Config
}

func NewConfigRepository(configPath string) ConfigRepository {
	return &ConfigRepositoryImpl{
		configPath:     configPath,
		configInstance: nil,
	}
}

func (c *ConfigRepositoryImpl) Get() *entity.Config {
	// If not initialized, load config
	if c.configInstance == nil {
		c.configInstance, _ = c.Load()
	}

	return c.configInstance
}

func (c *ConfigRepositoryImpl) Load() (config *entity.Config, err error) {
	config = &entity.Config{}

	// Read specific files
	jsonBytes, err := ioutil.ReadFile(c.configPath)
	if err != nil {
		return nil, err
	}

	// Unmarshal read json
	err = json.Unmarshal(jsonBytes, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}