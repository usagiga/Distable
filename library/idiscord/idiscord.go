package idiscord

import (
	"errors"
	"fmt"
	"github.com/vincent-petithory/dataurl"
	"io/ioutil"
	"net/http"
)

const (
	baseUri = "https://cdn.discordapp.com/"
	customEmojiEndpoint = "emojis/"
)

type APIClient struct {}

func NewIdiscord() *APIClient {
	return &APIClient{}
}

func (client *APIClient) GetEmoji(emojiId string, ext string) (imageUri *dataurl.DataURL, err error) {
	// Validate args
	if !(ext == "png" || ext == "gif") {
		errMsg := fmt.Sprintf("Extension \"%s\" is not compatible with this EP. You must use \"png\" or \"gif\"", ext)
		return nil, errors.New(errMsg)
	}

	// Get
	url := baseUri + customEmojiEndpoint + emojiId + "." + ext
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	// Convert to URI
	imgBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, nil
	}

	mediaType := "image/" + ext
	imageUri = dataurl.New(imgBytes, mediaType)

	return imageUri, nil
}