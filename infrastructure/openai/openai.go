package openai

import (
	"anki-support/helper"
	openai "github.com/sashabaranov/go-openai"
)

type OpenAI struct {
	openai *openai.Client
}

func NewClient() OpenAIer {
	client := openai.NewClient(helper.Config.OpenAIToken())
	return &OpenAI{openai: client}
}

//go:generate mockgen -destination=openai.mock.go -package=openai -self_package=anki-support/infrastructure/openai . OpenAIer
type OpenAIer interface {
	MakeJapaneseSentence(rememberVocabularyList []string, vocabulary, meaning string) (japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence string, err error)
}
