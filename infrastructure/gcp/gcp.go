package gcp

import (
	"anki-support/helper"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	log "github.com/sirupsen/logrus"
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
			time.Sleep(3 * time.Second)
			log.Warnf("fail to connect to google api, retry...")
			continue
		}
		return c
	}
}
func (c *Client) Close() {
	c.textToSpeechClient.Close()
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
