package main

import (
	"github.com/usagiga/distable/application"
	"github.com/usagiga/distable/domain"
	"github.com/usagiga/distable/infrastructure"
	"github.com/usagiga/distable/presentation"
	"github.com/usagiga/distable/repository"
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
