package application

import (
	"anki-support/domain"
	"fmt"
)

type AnkiNormalJapaneseOperator struct {
	originNote             domain.AnkiNote
	updateAudio            []domain.AnkiAudio
	noteFields             ankiNormalNoteField
	gpter                  domain.GPTer
	textToSpeecher         domain.TextToSpeecher
	ankier                 domain.Ankier
	rememberVocabularyList []string
}

type ankiNormalNoteField struct {
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
}

func (f ankiNormalNoteField) FieldDataMap() map[string]domain.AnkiFieldData {
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
	}
}

func NewAnkiNormalOperator(note domain.AnkiNote, gpter domain.GPTer, textToSpeecher domain.TextToSpeecher, ankier domain.Ankier, rememberVocabularyList []string) AnkiOperator {

	noteFields := ankiNormalNoteField{
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
	}
	return &AnkiNormalJapaneseOperator{
		originNote:             note,
		noteFields:             noteFields,
		gpter:                  gpter,
		textToSpeecher:         textToSpeecher,
		ankier:                 ankier,
		rememberVocabularyList: rememberVocabularyList,
	}
}

func (n *AnkiNormalJapaneseOperator) Do() error {
	_ = n.expressToSound()
	_ = n.expressToSentenceAndSound()
	n.originNote.Fields = n.noteFields.FieldDataMap()
	_ = n.ankier.UpdateNoteById(n.originNote.Id, n.originNote, n.updateAudio)
	n.ankier.AddNoteTagFromNoteId(n.originNote.Id, domain.AnkiDoneTagName)
	n.ankier.DeleteNoteTagFromNoteId(n.originNote.Id, domain.AnkiTodoTagName)
	return nil
}

func (n *AnkiNormalJapaneseOperator) expressToSound() error {
	if n.noteFields.japaneseToSound != "" {
		return nil
	}
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(n.noteFields.expression)
	n.noteFields.japaneseToSound = fmt.Sprintf("[sound:%s.mp3]", n.noteFields.expression)
	n.updateAudio = append(n.updateAudio, domain.AnkiAudio{
		Path:     expressFilePath,
		Filename: fmt.Sprintf("%s.mp3", n.noteFields.expression),
		Fields:   []string{"Japanese-ToSound"},
	})
	return nil
}

func (n *AnkiNormalJapaneseOperator) expressToSentenceAndSound() error {
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
	n.updateAudio = append(n.updateAudio, domain.AnkiAudio{
		Path:     sentenceFilePath,
		Filename: fmt.Sprintf("%s.mp3", sentence),
		Fields:   []string{"JapaneseSentence-ToSound"},
	})
	return nil
}
