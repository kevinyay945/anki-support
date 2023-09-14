package gcp

import (
	"anki-support/helper"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"testing"
)

type TextToSpeechSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	client   *Client
}

func TestSuiteInitTextToSpeech(t *testing.T) {
	if os.Getenv("RUN_INFRASTRUCTURE") == "true" {
		t.Skip("Skipping testing in production")
	}
	suite.Run(t, new(TextToSpeechSuite))
}

func (t *TextToSpeechSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	mockConfiger := helper.NewMockConfiger(t.mockCtrl)
	configData, _ := os.ReadFile("../../.config.dev.yaml")
	credential := struct {
		Token string `yaml:"GOOGLE_API_TOKEN"`
	}{}
	err := yaml.Unmarshal(configData, &credential)
	t.NoError(err, "Fail to parse config file")
	t.NotEmpty(credential.Token, "GCP Token is empty")

	mockConfiger.EXPECT().GoogleApiToken().Return(string(credential.Token))
	helper.Config = mockConfiger

	t.client = NewClient()
}

func (t *TextToSpeechSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *TextToSpeechSuite) Test_generate_japanese_file() {
	outputPath := "/Users/kevin/Developer/side-project/anki-support/temp"
	expectOutput := filepath.Join(outputPath, "私の机は木製です。.mp3")
	os.Remove(expectOutput)
	_, err := t.client.GenerateAudioByText("私の机は木製です。", outputPath, "私の机は木製です。")
	t.NoError(err)
	t.FileExists(expectOutput)
}
