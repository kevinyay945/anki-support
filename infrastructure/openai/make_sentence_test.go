package openai

import (
	"anki-support/helper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gopkg.in/yaml.v3"
	"os"
	"testing"
)

type MakeSentenceSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	client   *Client
}

func TestSuiteInitMakeSentence(t *testing.T) {
	if os.Getenv("RUN_INFRASTRUCTURE") == "true" {
		t.Skip("Skipping testing in production")
	}
	suite.Run(t, new(MakeSentenceSuite))
}

func (t *MakeSentenceSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	mockConfiger := helper.NewMockConfiger(t.mockCtrl)
	configData, _ := os.ReadFile("../../.config.dev.yaml")
	credential := struct {
		Token string `yaml:"OPEN_AI_TOKEN"`
	}{}
	err := yaml.Unmarshal(configData, &credential)
	t.NoError(err, "Fail to parse config file")
	t.NotEmpty(credential.Token, "OPEN AI Token is empty")
	mockConfiger.EXPECT().OpenAIToken().Return(string(credential.Token))
	helper.Config = mockConfiger
	t.client = NewClient()
}

func (t *MakeSentenceSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *MakeSentenceSuite) Test_make_japanese_sentence() {
	rememberVocabularyList := []string{
		"両親", "月餅", "電池", "彼の方", "お兄さん", "高速バス", "お姉さん", "映画", "前", "乗り場", "冷蔵庫", "学校", "明後日", "チケット", "番線",
	}
	sentence, hiraganaSentence, chineseSentence, err := t.client.makeJapaneseSentence(rememberVocabularyList, "山")
	t.NoError(err)

	t.Equal("japanese origin sentence", sentence)
	t.Equal("japanese hiragana sentence", hiraganaSentence)
	t.Equal("traditional Chinese sentence", chineseSentence)

}
