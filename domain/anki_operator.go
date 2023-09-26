package domain

import "fmt"

type OperatorGenerate struct {
	gpter          GPTer
	textToSpeecher TextToSpeecher
	ankier         Ankier
}

func NewOperatorGenerate(gpter GPTer, textToSpeecher TextToSpeecher, ankier Ankier) OperatorGenerator {
	return &OperatorGenerate{gpter: gpter, textToSpeecher: textToSpeecher, ankier: ankier}
}

func (g *OperatorGenerate) GetByNote(note AnkiNote, rememberVocabularyList []string) (o Operator, err error) {
	switch note.ModelName {
	case "Japanese (recognition&recall) 動詞篇":
		o = &VerbOperator{
			Note:                   note,
			gpter:                  g.gpter,
			textToSpeecher:         g.textToSpeecher,
			ankier:                 g.ankier,
			rememberVocabularyList: rememberVocabularyList,
		}
	case "Japanese (recognition&recall) 形容詞":
		o = &AdjOperator{
			Note:                   note,
			gpter:                  g.gpter,
			textToSpeecher:         g.textToSpeecher,
			ankier:                 g.ankier,
			rememberVocabularyList: rememberVocabularyList,
		}
	case "Japanese (recognition&recall)":
		o = NewNormalOperator(note, g.gpter, g.textToSpeecher, g.ankier, rememberVocabularyList)
	default:
		err = fmt.Errorf("don't support for this modelType: %s", note.ModelName)
	}
	return
}

type OperatorGenerator interface {
	GetByNote(note AnkiNote, rememberVoculary []string) (o Operator, err error)
}
