package domain

import (
	"fmt"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type OperatorSuite struct {
	suite.Suite
	mockCtrl           *gomock.Controller
	generator          OperatorGenerator
	mockGPTer          *MockGPTer
	mockAnkier         *MockAnkier
	mockTextToSpeecher *MockTextToSpeecher
}

func TestSuiteInitOperator(t *testing.T) {
	suite.Run(t, new(OperatorSuite))
}

func (t *OperatorSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.mockGPTer = NewMockGPTer(t.mockCtrl)
	t.mockAnkier = NewMockAnkier(t.mockCtrl)
	t.mockTextToSpeecher = NewMockTextToSpeecher(t.mockCtrl)
	t.generator = NewOperatorGenerate(t.mockGPTer, t.mockTextToSpeecher, t.mockAnkier)
}

func (t *OperatorSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *OperatorSuite) Test_normal_operator() {
	fields := map[string]FieldData{
		"Expression":                 {Value: "test expression value"},
		"Meaning":                    {Value: "test meaning value"},
		"Reading":                    {Value: "test reading value"},
		"Japanese-ToSound":           {Value: "test japanese to sound value"},
		"JapaneseSentence":           {Value: "test japanese hiragana sentence value"},
		"JapaneseSentence-ToSound":   {Value: "test japanese sentence to sound value"},
		"JapaneseSentence-ToChinese": {Value: "test japanese sentence to chinese value"},
		"Japanese-Note":              {Value: "test japanese note value"},
		"Japanese-ToChineseNote":     {Value: "test japanese to chinese note value"},
		"Answer-Note":                {Value: "test answer note value"},
	}
	note := AnkiNote{
		Id:        123,
		ModelName: "Japanese (recognition&recall)",
		Fields:    fields,
		Tags:      []string{"anki-helper-vocabulary-todo"},
	}
	japaneseSentence := "test japanese sentence"
	rememberVocabularyList := []string{"vocabulary1", "vocabulary2"}
	expressionVoicePath := "expression/file/path.mp3"
	japaneseSentenceVoicePath := "japanese/sentence/file/path.mp3"
	// expression to Japanese-ToSound by text to speech
	t.mockTextToSpeecher.EXPECT().
		GetJapaneseSound(fields["Expression"].Value).
		Return(expressionVoicePath, nil)
	// expression to JapaneseSentence and JapaneseSentence-ToChinese and JapaneseSentence-ToSound.Name by gpt
	t.mockGPTer.EXPECT().
		MakeJapaneseSentence(fields["Expression"].Value, fields["Meaning"].Value, rememberVocabularyList).
		Return(japaneseSentence, fields["JapaneseSentence"].Value, fields["JapaneseSentence-ToChinese"].Value, nil)
	// JapaneseSentence to JapaneseSentence-ToSound by text to speech
	t.mockTextToSpeecher.EXPECT().
		GetJapaneseSound(japaneseSentence).
		Return(japaneseSentenceVoicePath, nil)
	// ankier update note by id with expect data
	t.mockAnkier.EXPECT().
		UpdateNoteById(note.Id, note, []Audio{
			{
				Path:     expressionVoicePath,
				Filename: fmt.Sprintf("[Sound:%s.mp3]", fields["Meaning"].Value),
				Fields:   []string{"Japanese-ToSound"},
			},
			{
				Path:     japaneseSentenceVoicePath,
				Filename: fmt.Sprintf("[Sound:%s.mp3]", japaneseSentence),
				Fields:   []string{"JapaneseSentence-ToSound"},
			},
		}).
		Return(nil)
	// delete t-o-d-o tag in anki card
	t.mockAnkier.EXPECT().DeleteNoteTagFromNoteId(note.Id, AnkiTodoTagName).Return(nil)
	// add done tag at anki card
	t.mockAnkier.EXPECT().AddNoteTagFromNoteId(note.Id, AnkiDoneTagName).Return(nil)

	operator, err := t.generator.GetByNote(AnkiNote{
		Id:        note.Id,
		ModelName: note.ModelName,
		Fields: map[string]FieldData{
			"Expression":                 fields["Expression"],
			"Meaning":                    fields["Meaning"],
			"Reading":                    fields["Reading"],
			"Japanese-ToSound":           {Value: ""},
			"JapaneseSentence":           {Value: ""},
			"JapaneseSentence-ToSound":   {Value: ""},
			"JapaneseSentence-ToChinese": {Value: ""},
			"Japanese-Note":              {Value: "test japanese note value"},
			"Japanese-ToChineseNote":     {Value: "test japanese to chinese note value"},
			"Answer-Note":                {Value: "test answer note value"},
		},
		Tags: []string{"anki-helper-vocabulary-todo"},
	})
	t.NoError(err)
	err = operator.Do()
	t.NoError(err)
}
