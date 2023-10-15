package application

import "anki-support/domain"

type AnkiVerbOperator struct {
	Note                   domain.AnkiNote
	gpter                  domain.GPTer
	textToSpeecher         domain.TextToSpeecher
	ankier                 domain.Ankier
	rememberVocabularyList []string
}

func (v *AnkiVerbOperator) Do() error {
	//TODO implement me
	panic("implement me")
}
