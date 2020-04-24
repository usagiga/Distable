package presentation

import (
	"context"
	"github.com/usagiga/distable/application"
	"github.com/usagiga/distable/entity"
	"log"
	"os"
	"os/signal"
)

// EmojiSyncStreamCommandImpl is struct implemented `EmojiSyncStreamCommand`.
type EmojiSyncStreamCommandImpl struct {
	emojiSyncStreamApplication application.EmojiSyncStreamApplication
}

// NewEmojiSyncStreamCommand initializes `EmojiSyncStreamCommand`
func NewEmojiSyncStreamCommand(
	emojiSyncStreamApplication application.EmojiSyncStreamApplication,
) EmojiSyncStreamCommand {
	return &EmojiSyncStreamCommandImpl{
		emojiSyncStreamApplication: emojiSyncStreamApplication,
	}
}

// RunSyncService lets Distable wait updating emojis to sync them among specified servers.
func (e *EmojiSyncStreamCommandImpl) RunSyncService(servers []entity.ServerContext) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	errCh := make(chan error)

	// Run
	go func(){
		err := e.emojiSyncStreamApplication.WatchUpdatingEmoji(ctx, servers)
		if err != nil {
			errCh <- err
		}
	}()

	// Wait ^C or to raise errors
	interrupt := make(chan os.Signal)
	signal.Notify(interrupt, os.Interrupt)
	log.Println("Distable(Stream-Mode) is now running... Press ^C to quit")

	select {
	case <-interrupt:
		log.Println("Keyboard Interrupt is detected. Quiting...")
		cancel()
	case err := <-errCh:
		log.Fatal("Error raised: ", err)
	}
}
