package domain

import "anki-support/infrastructure/openai"

type GPT struct {
	openAIer openai.OpenAIer
}

func NewGPT(openAIer openai.OpenAIer) GPTer {
	return &GPT{openAIer: openAIer}
}

func (g *GPT) MakeJapaneseSentence(vocabulary string, meaning string, rememberVocabularyList []string) (string, string, string, error) {
	var err error
	sentence, hiraganaSentence, chineseSentence, err := g.openAIer.MakeJapaneseSentence(rememberVocabularyList, vocabulary, meaning)
	return sentence, hiraganaSentence, chineseSentence, err
}

type GPTer interface {
	MakeJapaneseSentence(vocabulary string, meaning string, rememberVocabularyList []string) (string, string, string, error)
}
