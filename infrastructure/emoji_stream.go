package infrastructure

import (
	"context"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/distable/entity"
)

// EmojiStreamInfraImpl is struct implemented `EmojiStreamInfra`.
type EmojiStreamInfraImpl struct {
	cred *entity.Credential
}

// NewEmojiStreamInfra initializes `EmojiStreamInfra`.
func NewEmojiStreamInfra(cred *entity.Credential) EmojiStreamInfra {
	return &EmojiStreamInfraImpl{
		cred: cred,
	}
}

// WatchUpdatingEmoji waits updating emoji on specified servers
// and notifies about it to the handler.
func (e *EmojiStreamInfraImpl) WatchUpdatingEmoji(
	ctx context.Context,
	onUpdatedEmojiHandler entity.OnUpdatedEmojiHandler,
) (err error) {
	errCh := make(chan error)
	token := e.cred.GetBearerToken()
	onUpdateEmoji := func(s *discordgo.Session, m *discordgo.GuildEmojisUpdate) {
		tgtServ := &entity.ServerContext{
			GuildID: m.GuildID,
		}

		// Fire OnUpdatedEmojiHandler
		err := onUpdatedEmojiHandler(tgtServ)
		if err != nil {
			errCh <- err
		}
	}

	// Listen the registered event.
	discord, err := discordgo.New(token)
	if err != nil {
		errCh <- err
		return
	}

	discord.AddHandler(onUpdateEmoji)
	err = discord.Open()
	if err != nil {
		errCh <- err
		return
	}
	defer discord.Close()

	// Wait for interrupting or raising errors.
	select {
	case err = <-errCh:
		return err
	case <-ctx.Done():
		return nil
	}
}
