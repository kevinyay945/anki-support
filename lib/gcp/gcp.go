package gcp

import (
	"anki-support/helper"
	texttospeech "cloud.google.com/go/texttospeech/apiv1"
	"context"
	log "github.com/sirupsen/logrus"
	"google.golang.org/api/option"
	"time"
)

type GCP struct {
	textToSpeechClient *texttospeech.Client
}

func NewGCP() GCPer {
	c := &GCP{}
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
func (c *GCP) Close() {
	_ = c.textToSpeechClient.Close()
}

func (c *GCP) setClientByToken(token string) (err error) {
	tokenByte := []byte(token)
	client, err := texttospeech.NewClient(context.Background(), option.WithCredentialsJSON(tokenByte))
	if err != nil {
		return
	}
	c.textToSpeechClient = client
	return
}

//go:generate mockgen -destination=gcp.mock.go -typed=true -package=gcp -self_package=anki-support/lib/gcp . GCPer
type GCPer interface {
	Close()
	setClientByToken(token string) (err error)
	GenerateAudioByText(inputText string, outputPath string, outputFileName string) (outputFilePath string, err error)
}
