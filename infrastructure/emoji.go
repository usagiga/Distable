package infrastructure

import (
	"errors"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/usagiga/Distable/entity"
	"github.com/usagiga/Distable/library/idiscord"
)

type EmojiInfraImpl struct {}

func NewEmojiInfra() EmojiInfra {
	return &EmojiInfraImpl{}
}

func (e *EmojiInfraImpl) Add(emoji *entity.Emoji, destServCtx *entity.ServerContext) (err error) {
	token := destServCtx.GetBearerToken()
	name := emoji.Name
	imgURI := emoji.ToURIString()

	// Create guild emoji
	discord, err := discordgo.New(token)
	if err != nil {
		return err
	}

	_, err = discord.GuildEmojiCreate(destServCtx.GuildID, name, imgURI, nil)

	return err
}

func (e *EmojiInfraImpl) Fetch(emojiCtx *entity.EmojiContext) (emoji *entity.Emoji, err error) {
	id := emojiCtx.ID
	ext := emojiCtx.GetExtension()

	// Fetch image
	imgClient := idiscord.NewIdiscord()
	imgUri, err := imgClient.GetEmoji(id, ext)
	if err != nil {
		return nil, err
	}

	// Return result
	emoji = &entity.Emoji{
		EmojiContext: *emojiCtx,
		DataURI: imgUri,
	}

	return emoji, nil
}

func (e *EmojiInfraImpl) FetchAll(emojiCtxs []entity.EmojiContext) (emojis []entity.Emoji, err error) {
	imgClient := idiscord.NewIdiscord()
	emojis = []entity.Emoji{}

	for _, emojiCtx := range emojiCtxs {
		id := emojiCtx.ID
		ext := emojiCtx.GetExtension()

		// Fetch image
		imgUri, err := imgClient.GetEmoji(id, ext)
		if err != nil {
			return nil, err
		}

		// Pack result
		emoji := entity.Emoji{
			EmojiContext: emojiCtx,
			DataURI: imgUri,
		}

		emojis = append(emojis, emoji)
	}

	return emojis, nil
}

func (e *EmojiInfraImpl) FetchAllContext(servCtx entity.ServerContext) (emojis []entity.EmojiContext, err error) {
	emojis = []entity.EmojiContext{}
	guildID := servCtx.GuildID
	token := servCtx.GetBearerToken()

	// Fetch guild emojis
	discord, err := discordgo.New(token)
	if err != nil {
		errMsg := fmt.Sprintf("EmojiInfra.FetchAllContext(): Can't establish connection using specific access token: %s", err)
		return nil, errors.New(errMsg)
	}

	guild, err := discord.Guild(guildID)
	if err != nil {
		errMsg := fmt.Sprintf("EmojiInfra.FetchAllContext(): Can't get guild(ID: %s) status: %s", guildID, err)
		return nil, errors.New(errMsg)
	}

	rawEmojis := guild.Emojis

	// Parse into internal type
	for _, rawEmoji := range rawEmojis {
		eCtx := entity.EmojiContext{
			ID: rawEmoji.ID,
			Name: rawEmoji.Name,
			RequireColons: rawEmoji.RequireColons,
			Animated: rawEmoji.Animated,
		}

		emojis = append(emojis, eCtx)
	}

	return emojis, nil
}

