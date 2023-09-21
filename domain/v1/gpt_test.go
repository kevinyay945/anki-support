package domain

import (
	"anki-support/infrastructure/openai"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type GPTSuite struct {
	suite.Suite
	mockCtrl     *gomock.Controller
	mockOpenAIer *openai.MockOpenAIer
	gpt          GPTer
}

func TestSuiteInitGPT(t *testing.T) {
	suite.Run(t, new(GPTSuite))
}

func (t *GPTSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.mockOpenAIer = openai.NewMockOpenAIer(t.mockCtrl)
	t.gpt = NewGPT(t.mockOpenAIer)
}

func (t *GPTSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *GPTSuite) Test_make_japanese_sentence() {
	targetVocabulary := "targetVol"
	targetVocabularyMeaning := "target vocabulary meaning"
	rememberVocabularyList := []string{"voc1", "voc2"}
	japaneseOriginSentence := "origin sentence"
	japaneseHiraganaSentence := "hiragana sentence"
	traditionalChineseSentence := "chinese sentence"

	var err error

	t.mockOpenAIer.EXPECT().
		MakeJapaneseSentence(rememberVocabularyList, targetVocabulary, targetVocabularyMeaning).
		Return(japaneseOriginSentence, japaneseHiraganaSentence, traditionalChineseSentence, err)

	japaneseSentence, hiraganaSentence, chineseTranslation, resErr := t.gpt.MakeJapaneseSentence(targetVocabulary, targetVocabularyMeaning, rememberVocabularyList)
	t.NoError(resErr)
	t.Equal(japaneseOriginSentence, japaneseSentence)
	t.Equal(japaneseHiraganaSentence, hiraganaSentence)
	t.Equal(traditionalChineseSentence, chineseTranslation)
}
