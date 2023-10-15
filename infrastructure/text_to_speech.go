package infrastructure

import (
	"anki-support/domain"
	"anki-support/helper"
	"anki-support/lib/gcp"
	"path/filepath"
)

type TextToSpeech struct {
	gcp gcp.GCPer
}

func NewTextToSpeech(gcp gcp.GCPer) domain.TextToSpeecher {
	return &TextToSpeech{gcp: gcp}
}

func (s *TextToSpeech) GetJapaneseSound(japaneseText string) (filePath string, err error) {
	path := helper.Config.AssetPath()
	outputFolder := filepath.Join(path, "japaneseSound")
	filePath, err = s.gcp.GenerateAudioByText(japaneseText, outputFolder, japaneseText)
	return
}
