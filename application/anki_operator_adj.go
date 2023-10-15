package application

import "anki-support/domain"

type AnkiAdjOperator struct {
	Note                   domain.AnkiNote
	gpter                  domain.GPTer
	textToSpeecher         domain.TextToSpeecher
	ankier                 domain.Ankier
	rememberVocabularyList []string
}

func (a *AnkiAdjOperator) Do() error {
	//TODO implement me
	panic("implement me")
}
