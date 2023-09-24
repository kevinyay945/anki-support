package domain

import (
	"anki-support/helper"
	"anki-support/infrastructure/gcp"
	"fmt"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"gopkg.in/yaml.v3"
	"os"
	"path/filepath"
	"testing"
)

type TextToSpeechSuite struct {
	suite.Suite
	mockCtrl   *gomock.Controller
	configPath string
}

func TestSuiteInitTextToSpeech(t *testing.T) {
	suite.Run(t, new(TextToSpeechSuite))
}

func (t *TextToSpeechSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.settingHelperConfiger()
}

func (t *TextToSpeechSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *TextToSpeechSuite) Test_text_to_japanese_speech() {
	gcpClient := gcp.NewMockGCPer(t.mockCtrl)
	speech := TextToSpeech{
		gcp: gcpClient,
	}
	convertJapaneseText := "これはペンです"
	outputFolder := filepath.Join(t.configPath, "/japaneseSound")
	outputFileName := fmt.Sprintf("%s.mp3", convertJapaneseText)
	gcpClient.EXPECT().GenerateAudioByText(
		convertJapaneseText,
		outputFolder,
		outputFileName,
	).Return(filepath.Join(outputFolder, outputFileName), nil)
	filePath, err := speech.GetJapaneseSound(
		convertJapaneseText,
	)
	t.NoError(err)
	t.Equal(filepath.Join(outputFolder, outputFileName), filePath)
}

func (t *TextToSpeechSuite) settingHelperConfiger() {
	mockConfiger := helper.NewMockConfiger(t.mockCtrl)
	configData, _ := os.ReadFile("../../.config.dev.yaml")
	config := struct {
		Path string `yaml:"ASSET_PATH"`
	}{}
	err := yaml.Unmarshal(configData, &config)
	t.NoError(err, "Fail to parse config file")
	t.NotEmpty(config.Path, "temp folder is empty")

	mockConfiger.EXPECT().AssetPath().Return(string(config.Path))
	helper.Config = mockConfiger
	t.configPath = config.Path
}
