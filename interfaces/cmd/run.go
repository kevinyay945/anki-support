package cmd

import (
	"anki-support/application"
	"errors"
)

type RunCmd struct {
	ankiRepository      application.AnkiRepositorier
	ankiOperatorFactory application.AnkiOperatorFactorier
}

func (c *RunCmd) RunForSpecificDeck(deckName string) (outputErr error) {
	allNoteList, err := c.ankiRepository.GetAllNotesByDeckName(deckName)
	if err != nil {
		outputErr = err
		return
	}
	var allVocabularyList []string
	for _, note := range allNoteList {
		vocabulary := note.Fields["Expression"].Value
		allVocabularyList = append(allVocabularyList, vocabulary)
	}
	todoNoteList, err := c.ankiRepository.GetAllTodoNotesByDeckName(deckName)
	if err != nil {
		outputErr = err
		return
	}
	for _, note := range todoNoteList {
		ankiOperator, err := c.ankiOperatorFactory.CreateByNote(note, allVocabularyList)
		if err != nil {
			outputErr = errors.Join(err)
			continue
		}
		err = ankiOperator.Do()
		if err != nil {
			outputErr = errors.Join(err)
			continue
		}
	}
	return
}
