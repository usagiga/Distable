package infrastructure

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/Distable/entity"
	"github.com/usagiga/Distable/library/idiscord"
)

// EmojiInfraImpl is struct implemented `EmojiInfra`.
type EmojiInfraImpl struct{}

// NewEmojiInfra initializes `EmojiInfra`.
func NewEmojiInfra() EmojiInfra {
	return &EmojiInfraImpl{}
}

// Add adds emoji into a specific server.
func (e *EmojiInfraImpl) Add(emoji *entity.Emoji, destServCtx *entity.ServerContext) (err error) {
	token := destServCtx.GetBearerToken()
	name := emoji.Name
	imgURI := emoji.ToURIString()

	// Establish a connection to the Discord API server.
	discord, err := discordgo.New(token)
	if err != nil {
		return err
	}

	// Create guild emoji.
	_, err = discord.GuildEmojiCreate(destServCtx.GuildID, name, imgURI, nil)

	return err
}

// Fetch fetches actual emoji image from its context.
func (e *EmojiInfraImpl) Fetch(emojiCtx *entity.EmojiContext) (emoji *entity.Emoji, err error) {
	id := emojiCtx.ID
	ext := emojiCtx.GetExtension()
	imgClient := idiscord.NewIdiscord()

	// Fetch image.
	imgUri, err := imgClient.GetEmoji(id, ext)
	if err != nil {
		return nil, err
	}

	// Parse results into internal type.
	emoji = &entity.Emoji{
		EmojiContext: *emojiCtx,
		DataURI:      imgUri,
	}

	return emoji, nil
}

// FetchAll fetches actual emoji image from its context array one-by-one.
func (e *EmojiInfraImpl) FetchAll(emojiCtxs []entity.EmojiContext) (emojis []entity.Emoji, err error) {
	imgClient := idiscord.NewIdiscord()
	emojis = []entity.Emoji{}

	for _, emojiCtx := range emojiCtxs {
		id := emojiCtx.ID
		ext := emojiCtx.GetExtension()

		// Fetch image.
		imgUri, err := imgClient.GetEmoji(id, ext)
		if err != nil {
			return nil, err
		}

		// Parse results into internal type.
		emoji := entity.Emoji{
			EmojiContext: emojiCtx,
			DataURI:      imgUri,
		}

		emojis = append(emojis, emoji)
	}

	return emojis, nil
}

// FetchAllContext fetches emoji contexts on a specific server.
func (e *EmojiInfraImpl) FetchAllContext(servCtx entity.ServerContext) (emojis []entity.EmojiContext, err error) {
	emojis = []entity.EmojiContext{}
	guildID := servCtx.GuildID
	token := servCtx.GetBearerToken()

	// Establish a connection to the Discord API server.
	discord, err := discordgo.New(token)
	if err != nil {
		errMsg := fmt.Sprintf("EmojiInfra.FetchAllContext(): Can't establish connection using specific access token: %s", err)
		return nil, errors.New(errMsg)
	}

	// Fetch guild emojis.
	guild, err := discord.Guild(guildID)
	if err != nil {
		errMsg := fmt.Sprintf("EmojiInfra.FetchAllContext(): Can't get guild(ID: %s) status: %s", guildID, err)
		return nil, errors.New(errMsg)
	}

	rawEmojis := guild.Emojis

	// Parse results into internal type.
	for _, rawEmoji := range rawEmojis {
		eCtx := entity.EmojiContext{
			ID:            rawEmoji.ID,
			Name:          rawEmoji.Name,
			RequireColons: rawEmoji.RequireColons,
			Animated:      rawEmoji.Animated,
		}

		emojis = append(emojis, eCtx)
	}

	return emojis, nil
}
