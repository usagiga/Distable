package main

import (
	"github.com/usagiga/Distable/application"
	"github.com/usagiga/Distable/domain"
	"github.com/usagiga/Distable/infrastructure"
	"github.com/usagiga/Distable/presentation"
	"github.com/usagiga/Distable/repository"
)

const (
	configPath = "./config.json"
)

func main() {
	// Config
	configRepos := repository.NewConfigRepository(configPath)
	config := configRepos.Get()

	// Model
	emojiArrayModel := domain.NewEmojiArrayModel()

	// Repository
	// None

	// Infrastructure
	emojiInfra := infrastructure.NewEmojiInfra()

	// Application
	emojiSyncApplication := application.NewEmojiSyncApplication(emojiArrayModel, emojiInfra)

	// Presentation
	emojiSyncCommand := presentation.NewEmojiSyncCommand(emojiSyncApplication)

	// Run app
	servers := config.Servers
	emojiSyncCommand.Sync(servers)
}
