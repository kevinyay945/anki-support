package domain

import (
	"anki-support/helper"
	"anki-support/infrastructure/gcp"
	"fmt"
	"path/filepath"
)

type TextToSpeech struct {
	gcp gcp.GCPer
}

func (s *TextToSpeech) JapaneseSound(japaneseText string) (filePath string, err error) {
	path := helper.Config.AssetPath()
	outputFolder := filepath.Join(path, "japaneseSound")
	filePath, err = s.gcp.GenerateAudioByText(japaneseText, outputFolder, fmt.Sprintf("%s.mp3", japaneseText))
	return
}

type TextToSpeecher interface {
	JapaneseSound(japaneseText string) (filePath string, err error)
}
