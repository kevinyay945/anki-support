package gcp

import (
	"anki-support/helper"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	"google.golang.org/api/option"
	"time"
)

type Client struct {
	textToSpeechClient *texttospeech.Client
}

func NewClient() *Client {
	c := &Client{}
	for {
		token := helper.Config.GoogleApiToken()
		if err := c.setClientByToken(token); err != nil {
			continue
		}
		time.Sleep(3 * time.Second)
		return c
	}
}

func (c *Client) setClientByToken(token string) (err error) {
	tokenByte := []byte(token)
	client, err := texttospeech.NewClient(context.Background(), option.WithCredentialsJSON(tokenByte))
	if err != nil {
		return
	}
	c.textToSpeechClient = client
	return
}
