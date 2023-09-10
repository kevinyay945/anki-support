package openai

import (
	"anki-support/helper"
	openai "github.com/sashabaranov/go-openai"
)

type Client struct {
	openai *openai.Client
}

func NewClient() *Client {
	client := openai.NewClient(helper.Config.OpenAIToken())
	return &Client{openai: client}
}
