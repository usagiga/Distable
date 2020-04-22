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
	// Parse args
	useStream := ParseArgs()

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
	emojiStreamInfra := infrastructure.NewEmojiStreamInfra(cred)

	// Application
	emojiSyncApplication := application.NewEmojiSyncApplication(emojiArrayModel, serverArrayModel, emojiInfra)
	emojiSyncStreamApplication := application.NewEmojiSyncStreamApplication(emojiSyncApplication, emojiStreamInfra)

	// Presentation
	emojiSyncCommand := presentation.NewEmojiSyncCommand(emojiSyncApplication)
	emojiSyncStreamCommand := presentation.NewEmojiSyncStreamCommand(emojiSyncStreamApplication)

	// Run app
	switch {
	case useStream:
		emojiSyncStreamCommand.RunSyncService(servers)
	default:
		emojiSyncCommand.Sync(servers)
	}
}
