package domain

import "fmt"

type Operator interface {
	Do() error
}

type NormalJapaneseOperator struct {
	originNote AnkiNote
	noteFields struct {
		express                   string
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
	gpter                  GPTer
	textToSpeecher         TextToSpeecher
	ankier                 Ankier
	rememberVocabularyList []string
}

type normalNoteField struct {
	express                   string
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

func (f normalNoteField) FieldDataMap() map[string]FieldData {
	return map[string]FieldData{
		"Expression":                 {f.express, 0},
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

func NewNormalOperator(note AnkiNote, gpter GPTer, textToSpeecher TextToSpeecher, ankier Ankier, rememberVocabularyList []string) *NormalJapaneseOperator {

	noteFields := normalNoteField{
		express:                   note.Fields["Expression"].Value,
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
	return &NormalJapaneseOperator{
		originNote:             note,
		noteFields:             noteFields,
		gpter:                  gpter,
		textToSpeecher:         textToSpeecher,
		ankier:                 ankier,
		rememberVocabularyList: rememberVocabularyList,
	}
}

func (n *NormalJapaneseOperator) Do() error {
	expressFilePath, _ := n.textToSpeecher.GetJapaneseSound(n.noteFields.express)
	sentence, hiraganaSentence, chineseSentence, _ := n.gpter.MakeJapaneseSentence(n.noteFields.express, n.noteFields.meaning, n.rememberVocabularyList)
	sentenceFilePath, _ := n.textToSpeecher.GetJapaneseSound(sentence)
	field := normalNoteField{
		express:                   n.noteFields.express,
		meaning:                   n.noteFields.meaning,
		reading:                   n.noteFields.reading,
		japaneseToSound:           fmt.Sprintf("[sound:%s.mp3]", n.noteFields.express),
		japaneseSentence:          hiraganaSentence,
		japaneseSentenceToSound:   fmt.Sprintf("[sound:%s.mp3]", sentence),
		japaneseSentenceToChinese: chineseSentence,
		japaneseNote:              n.noteFields.japaneseNote,
		japaneseToChineseNote:     n.noteFields.japaneseToChineseNote,
		answerNote:                n.noteFields.answerNote,
	}
	n.originNote.Fields = field.FieldDataMap()
	_ = n.ankier.UpdateNoteById(n.originNote.Id, n.originNote, []Audio{
		{
			Path:     expressFilePath,
			Filename: fmt.Sprintf("%s.mp3", n.noteFields.express),
			Fields:   []string{"Japanese-ToSound"},
		},
		{
			Path:     sentenceFilePath,
			Filename: fmt.Sprintf("%s.mp3", sentence),
			Fields:   []string{"JapaneseSentence-ToSound"},
		},
	})
	n.ankier.AddNoteTagFromNoteId(n.originNote.Id, AnkiDoneTagName)
	n.ankier.DeleteNoteTagFromNoteId(n.originNote.Id, AnkiTodoTagName)
	return nil
}

type VerbOperator struct {
	Note                   AnkiNote
	gpter                  GPTer
	textToSpeecher         TextToSpeecher
	ankier                 Ankier
	rememberVocabularyList []string
}

func (v *VerbOperator) Do() error {
	//TODO implement me
	panic("implement me")
}

type AdjOperator struct {
	Note                   AnkiNote
	gpter                  GPTer
	textToSpeecher         TextToSpeecher
	ankier                 Ankier
	rememberVocabularyList []string
}

func (a *AdjOperator) Do() error {
	//TODO implement me
	panic("implement me")
}
