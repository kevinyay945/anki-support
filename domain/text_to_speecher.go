package domain

//go:generate mockgen -destination=text_to_speecher.mock.go -typed=true -package=domain -self_package=anki-support/domain . TextToSpeecher
type TextToSpeecher interface {
	GetJapaneseSound(japaneseText string) (filePath string, err error)
}
