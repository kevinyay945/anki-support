package application

import (
	"anki-support/domain"
	"fmt"
)

type ankiAdjNoteField struct {
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
	pastNegative              string
	pastNegativeToSound       string
}

func (f ankiAdjNoteField) FieldDataMap() map[string]domain.AnkiFieldData {
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
		"PastNegative":               {f.pastNegative, 16},
		"PastNegative-ToSound":       {f.pastNegativeToSound, 17},
	}
}

type AnkiAdjOperator struct {
	gpter                  domain.GPTer
	textToSpeecher         domain.TextToSpeecher
	ankier                 domain.Ankier
	rememberVocabularyList []string
	noteFields             ankiAdjNoteField
	originNote             domain.AnkiNote
	updateAudioList        []domain.AnkiAudio
}

func NewAnkiAdjOperator(gpter domain.GPTer, textToSpeecher domain.TextToSpeecher, ankier domain.Ankier, rememberVocabularyList []string, note domain.AnkiNote) AnkiOperator {
	noteFields := ankiAdjNoteField{
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
		pastNegative:              note.Fields["PastNegative"].Value,
		pastNegativeToSound:       note.Fields["PastNegative-ToSound"].Value,
	}
	return &AnkiAdjOperator{
		gpter:                  gpter,
		textToSpeecher:         textToSpeecher,
		ankier:                 ankier,
		rememberVocabularyList: rememberVocabularyList,
		originNote:             note,
		updateAudioList:        []domain.AnkiAudio{},
		noteFields:             noteFields,
	}
}

func (n *AnkiAdjOperator) Do() error {
	_ = n.expressToSound()
	_ = n.expressToSentenceAndSound()
	_ = n.originalToSound()
	_ = n.negativeToSound()
	_ = n.pastToSound()
	_ = n.pastNegativeToSound()

	n.originNote.Fields = n.noteFields.FieldDataMap()
	_ = n.ankier.UpdateNoteById(n.originNote.Id, n.originNote, n.updateAudioList)
	n.ankier.AddNoteTagFromNoteId(n.originNote.Id, domain.AnkiDoneTagName)
	n.ankier.DeleteNoteTagFromNoteId(n.originNote.Id, domain.AnkiTodoTagName)
	return nil
}

func (n *AnkiAdjOperator) expressToSound() error {
	if n.noteFields.japaneseToSound != "" {
		return nil
	}
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(n.noteFields.expression)
	n.noteFields.japaneseToSound = fmt.Sprintf("[sound:%s.mp3]", n.noteFields.expression)
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", n.noteFields.expression),
		Fields:   []string{"Japanese-ToSound"},
	})
	return nil
}

func (n *AnkiAdjOperator) expressToSentenceAndSound() error {
	if n.noteFields.japaneseSentenceToSound != "" {
		return nil
	}

	var sentence, hiraganaSentence, chineseSentence string
	if n.noteFields.japaneseSentence == "" {
		sentence, hiraganaSentence, chineseSentence, _ = n.gpter.MakeJapaneseSentence(n.noteFields.expression, n.noteFields.meaning, n.rememberVocabularyList)
	} else {
		sentence = n.noteFields.japaneseSentence
		hiraganaSentence = n.noteFields.japaneseSentence
		chineseSentence = n.noteFields.japaneseSentenceToChinese
	}
	sentenceFilePath, _ := n.textToSpeecher.GetJapaneseSound(sentence)
	n.noteFields.japaneseSentence = hiraganaSentence
	n.noteFields.japaneseSentenceToSound = fmt.Sprintf("[sound:%s.mp3]", sentence)
	n.noteFields.japaneseSentenceToChinese = chineseSentence
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     sentenceFilePath,
		Filename: fmt.Sprintf("%s.mp3", sentence),
		Fields:   []string{"JapaneseSentence-ToSound"},
	})
	return nil
}

func (n *AnkiAdjOperator) originalToSound() error {
	if n.noteFields.originalToSound != "" || n.noteFields.original == "" {
		return nil
	}
	fileName := n.noteFields.original
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(fileName)
	n.noteFields.originalToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Original-ToSound"},
	})
	return nil
}

func (n *AnkiAdjOperator) negativeToSound() error {
	if n.noteFields.negativeToSound != "" || n.noteFields.negative == "" {
		return nil
	}
	fileName := n.noteFields.negative
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(fileName)
	n.noteFields.negativeToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Negative-ToSound"},
	})
	return nil
}

func (n *AnkiAdjOperator) pastToSound() error {
	if n.noteFields.pastToSound != "" || n.noteFields.past == "" {
		return nil
	}
	fileName := n.noteFields.past
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(fileName)
	n.noteFields.pastToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"Past-ToSound"},
	})
	return nil
}
func (n *AnkiAdjOperator) pastNegativeToSound() error {
	if n.noteFields.pastNegativeToSound != "" || n.noteFields.pastNegative == "" {
		return nil
	}
	fileName := n.noteFields.pastNegative
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(fileName)
	n.noteFields.pastNegativeToSound = fmt.Sprintf("[sound:%s.mp3]", fileName)
	n.updateAudioList = append(n.updateAudioList, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", fileName),
		Fields:   []string{"PastNegative-ToSound"},
	})
	return nil
}
