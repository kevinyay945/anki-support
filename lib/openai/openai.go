package openai

import (
	"anki-support/helper"
	openai "github.com/sashabaranov/go-openai"
	"github.com/sirupsen/logrus"
)

type OpenAI struct {
	openai *openai.Client
	log    *logrus.Entry
}

func NewClient(log *logrus.Logger) OpenAIer {

	client := openai.NewClient(helper.Config.OpenAIToken())
	return &OpenAI{
		openai: client,
		log:    log.WithField("package", "lib/openai"),
	}
}

//go:generate mockgen -destination=openai.mock.go -typed=true -package=openai -self_package=anki-support/lib/openai . OpenAIer
type OpenAIer interface {
	MakeJapaneseSentence(rememberVocabularyList []string, vocabulary, meaning string) (japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence string, err error)
}
