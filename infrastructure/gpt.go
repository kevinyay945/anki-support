package infrastructure

import (
	"anki-support/domain"
	"anki-support/lib/openai"
)

type GPT struct {
	openAIer openai.OpenAIer
}

func NewGPT(openAIer openai.OpenAIer) domain.GPTer {
	return &GPT{openAIer: openAIer}
}

func (g *GPT) MakeJapaneseSentence(vocabulary string, meaning string, rememberVocabularyList []string) (sentence string, hiraganaSentence string, chineseSentence string, err error) {
	sentence, hiraganaSentence, chineseSentence, err = g.openAIer.MakeJapaneseSentence(rememberVocabularyList, vocabulary, meaning)
	return sentence, hiraganaSentence, chineseSentence, err
}
