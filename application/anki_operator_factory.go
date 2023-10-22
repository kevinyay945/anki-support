package application

import (
	"anki-support/domain"
	"fmt"
)

type AnkiOperatorFactory struct {
	gpter          domain.GPTer
	textToSpeecher domain.TextToSpeecher
	ankier         domain.Ankier
}

func NewAnkiOperatorFactory(gpter domain.GPTer, textToSpeecher domain.TextToSpeecher, ankier domain.Ankier) AnkiOperatorFactorier {
	return &AnkiOperatorFactory{gpter: gpter, textToSpeecher: textToSpeecher, ankier: ankier}
}

func (g *AnkiOperatorFactory) CreateByNote(note domain.AnkiNote, rememberVocabularyList []string) (o AnkiOperator, err error) {
	switch note.ModelName {
	case "Japanese (recognition&recall) 動詞篇":
		o = NewAnkiVerbOperator(g.gpter, g.textToSpeecher, g.ankier, rememberVocabularyList, note)
	case "Japanese (recognition&recall) 形容詞":
		o = NewAnkiAdjOperator(g.gpter, g.textToSpeecher, g.ankier, rememberVocabularyList, note)
	case "Japanese (recognition&recall)":
		o = NewAnkiNormalOperator(note, g.gpter, g.textToSpeecher, g.ankier, rememberVocabularyList)
	default:
		err = fmt.Errorf("don't support for this modelType: %s", note.ModelName)
	}
	return
}

//go:generate mockgen -destination=anki_operator_factory.mock.go -typed=true -package=application -self_package=anki-support/application . AnkiOperatorFactorier
type AnkiOperatorFactorier interface {
	CreateByNote(note domain.AnkiNote, rememberVocabulary []string) (o AnkiOperator, err error)
}
