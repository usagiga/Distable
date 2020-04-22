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
	servers := config.Servers
	cred := config.Credential

	// Model
	emojiArrayModel := domain.NewEmojiArrayModel()
	serverArrayModel := domain.NewServerArrayModel()

	// Repository
	// None

	// Infrastructure
	emojiInfra := infrastructure.NewEmojiInfra(cred)

	// Application
	emojiSyncApplication := application.NewEmojiSyncApplication(emojiArrayModel, serverArrayModel, emojiInfra)

	// Presentation
	emojiSyncCommand := presentation.NewEmojiSyncCommand(emojiSyncApplication)

	// Run app
	emojiSyncCommand.Sync(servers)
}
