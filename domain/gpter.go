package domain

//go:generate mockgen -destination=gpter.mock.go -typed=true -package=domain -self_package=anki-support/domain . GPTer
type GPTer interface {
	MakeJapaneseSentence(vocabulary string, meaning string, rememberVocabularyList []string) (sentence string, hiraganaSentence string, chineseSentence string, err error)
}
