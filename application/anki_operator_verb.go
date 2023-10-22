package application

import (
	"anki-support/domain"
	"fmt"
)

type ankiVerbNoteField struct {
	expression                string
	meaning                   string
	reading                   string
	japaneseToSound           string
	japaneseSentence          string
	japaneseSentenceToSound   string
	japaneseSentenceToChinese string
	japaneseNote              string
	japaneseToChineseNote     string
	answerNote                string
	original                  string
	originalToSound           string
	negative                  string
	negativeToSound           string
	past                      string
	pastToSound               string
}

func (f ankiVerbNoteField) FieldDataMap() map[string]domain.AnkiFieldData {
	return map[string]domain.AnkiFieldData{
		"Expression":                 {f.expression, 0},
		"Meaning":                    {f.meaning, 1},
		"Reading":                    {f.reading, 2},
		"Japanese-ToSound":           {f.japaneseToSound, 3},
		"JapaneseSentence":           {f.japaneseSentence, 4},
		"JapaneseSentence-ToSound":   {f.japaneseSentenceToSound, 5},
		"JapaneseSentence-ToChinese": {f.japaneseSentenceToChinese, 6},
		"Japanese-Note":              {f.japaneseNote, 7},
		"Japanese-ToChineseNote":     {f.japaneseToChineseNote, 8},
		"Answer-Note":                {f.answerNote, 9},
		"Original":                   {f.original, 10},
		"Original-ToSound":           {f.originalToSound, 11},
		"Negative":                   {f.negative, 12},
		"Negative-ToSound":           {f.negativeToSound, 13},
		"Past":                       {f.past, 14},
		"Past-ToSound":               {f.pastToSound, 15},
	}
}

type AnkiVerbOperator struct {
	Note                   domain.AnkiNote
	gpter                  domain.GPTer
	textToSpeecher         domain.TextToSpeecher
	ankier                 domain.Ankier
	rememberVocabularyList []string
	noteFields             ankiVerbNoteField
	originNote             domain.AnkiNote
	updateAudioList        []domain.AnkiAudio
}

func NewAnkiVerbOperator(gpter domain.GPTer, textToSpeecher domain.TextToSpeecher, ankier domain.Ankier, rememberVocabularyList []string, note domain.AnkiNote) AnkiOperator {
	noteFields := ankiVerbNoteField{
		expression:                note.Fields["Expression"].Value,
		meaning:                   note.Fields["Meaning"].Value,
		reading:                   note.Fields["Reading"].Value,
		japaneseToSound:           note.Fields["Japanese-ToSound"].Value,
		japaneseSentence:          note.Fields["JapaneseSentence"].Value,
		japaneseSentenceToSound:   note.Fields["JapaneseSentence-ToSound"].Value,
		japaneseSentenceToChinese: note.Fields["JapaneseSentence-ToChinese"].Value,
		japaneseNote:              note.Fields["Japanese-Note"].Value,
		japaneseToChineseNote:     note.Fields["Japanese-ToChineseNote"].Value,
		answerNote:                note.Fields["Answer-Note"].Value,
		original:                  note.Fields["Original"].Value,
		originalToSound:           note.Fields["Original-ToSound"].Value,
		negative:                  note.Fields["Negative"].Value,
		negativeToSound:           note.Fields["Negative-ToSound"].Value,
		past:                      note.Fields["Past"].Value,
		pastToSound:               note.Fields["Past-ToSound"].Value,
	}
	return &AnkiVerbOperator{
		gpter:                  gpter,
		textToSpeecher:         textToSpeecher,
		ankier:                 ankier,
		rememberVocabularyList: rememberVocabularyList,
		originNote:             note,
		updateAudioList:        []domain.AnkiAudio{},
		noteFields:             noteFields,
	}
}
func (a *AnkiVerbOperator) Do() error {
	_ = a.expressToSound()
	_ = a.expressToSentenceAndSound()
	_ = a.originalToSound()
	_ = a.negativeToSound()
	_ = a.pastToSound()

	a.originNote.Fields = a.noteFields.FieldDataMap()
	_ = a.ankier.UpdateNoteById(a.originNote.Id, a.originNote, a.updateAudioList)
	a.ankier.AddNoteTagFromNoteId(a.originNote.Id, domain.AnkiDoneTagName)
	a.ankier.DeleteNoteTagFromNoteId(a.originNote.Id, domain.AnkiTodoTagName)
	return nil
}
func (a *AnkiVerbOperator) expressToSound() error {
	if a.noteFields.japaneseToSound != "" {
		return nil
	}
	expressFilePath, _ := a.textToSpeecher.GetJapaneseSound(a.noteFields.expression)
	a.noteFields.japaneseToSound = fmt.Sprintf("[sound:%s.mp3]", a.noteFields.expression)
	a.updateAudioList = append(a.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", a.noteFields.expression),
		Fields:   []string{"Japanese-ToSound"},
	})
	return nil
}

func (a *AnkiVerbOperator) expressToSentenceAndSound() error {
	if a.noteFields.japaneseSentenceToSound != "" {
		return nil
	}

	var sentence, hiraganaSentence, chineseSentence string
	if a.noteFields.japaneseSentence == "" {
		sentence, hiraganaSentence, chineseSentence, _ = a.gpter.MakeJapaneseSentence(a.noteFields.expression, a.noteFields.meaning, a.rememberVocabularyList)
	} else {
		sentence = a.noteFields.japaneseSentence
		hiraganaSentence = a.noteFields.japaneseSentence
		chineseSentence = a.noteFields.japaneseSentenceToChinese
	}
	sentenceFilePath, _ := a.textToSpeecher.GetJapaneseSound(sentence)
	a.noteFields.japaneseSentence = hiraganaSentence
	a.noteFields.japaneseSentenceToSound = fmt.Sprintf("[sound:%s.mp3]", sentence)
	a.noteFields.japaneseSentenceToChinese = chineseSentence
	a.updateAudioList = append(a.updateAudioList, domain.AnkiAudio{
		Path:     sentenceFilePath,
		Filename: fmt.Sprintf("%s.mp3", sentence),
		Fields:   []string{"JapaneseSentence-ToSound"},
	})
	return nil
}

func (a *AnkiVerbOperator) originalToSound() error {
	if a.noteFields.originalToSound != "" || a.noteFields.original == "" {
		return nil
	}
	fileName := a.noteFields.original
	expressFilePath, _ := a.textToSpeecher.GetJapaneseSound(fileName)
	a.noteFields.originalToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	a.updateAudioList = append(a.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Original-ToSound"},
	})
	return nil
}

func (a *AnkiVerbOperator) negativeToSound() error {
	if a.noteFields.negativeToSound != "" || a.noteFields.negative == "" {
		return nil
	}
	fileName := a.noteFields.negative
	expressFilePath, _ := a.textToSpeecher.GetJapaneseSound(fileName)
	a.noteFields.negativeToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	a.updateAudioList = append(a.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Negative-ToSound"},
	})
	return nil
}

func (a *AnkiVerbOperator) pastToSound() error {
	if a.noteFields.pastToSound != "" || a.noteFields.past == "" {
		return nil
	}
	fileName := a.noteFields.past
	expressFilePath, _ := a.textToSpeecher.GetJapaneseSound(fileName)
	a.noteFields.pastToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	a.updateAudioList = append(a.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Past-ToSound"},
	})
	return nil
}
