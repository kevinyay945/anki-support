package infrastructure

import (
	anki2 "anki-support/domain/anki"
	"anki-support/lib/anki"
	"github.com/atselvan/ankiconnect"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type AnkiSuite struct {
	suite.Suite
	mockCtrl   *gomock.Controller
	anki       anki2.Ankier
	mockAnkier *anki.MockAnkier
}

func TestSuiteInitAnki(t *testing.T) {
	suite.Run(t, new(AnkiSuite))
}

func (t *AnkiSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.mockAnkier = anki.NewMockAnkier(t.mockCtrl)
	t.anki = NewAnki(t.mockAnkier)
}

func (t *AnkiSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *AnkiSuite) Test_get_note_list_by_deck_name() {
	exampleNote, resultNotesInfo := t.getInfrastructureAndDomainNote()
	t.mockAnkier.EXPECT().GetAllNoteFromDeck("deckName").
		Return([]ankiconnect.ResultNotesInfo{resultNotesInfo}, nil)
	note, err := t.anki.GetNoteListByDeckName("deckName")
	t.Equal(nil, err)
	t.Equal([]anki2.Note{exampleNote}, note)
}

func (t *AnkiSuite) Test_get_note_by_id() {
	exampleNote, resultNotesInfo := t.getInfrastructureAndDomainNote()
	t.mockAnkier.EXPECT().GetNoteById(int64(123)).Return(resultNotesInfo, nil)
	note, err := t.anki.GetNoteById(123)
	t.Equal(nil, err)
	t.Equal(exampleNote, note)
}

func (t *AnkiSuite) Test_get_todo_noteFromDeck() {
	deckName := "deckName"
	domainNote, notesInfo := t.getInfrastructureAndDomainNote()
	t.mockAnkier.EXPECT().GetTodoNoteFromDeck(deckName).Return([]ankiconnect.ResultNotesInfo{notesInfo}, nil)
	noteList, err := t.anki.GetTodoNoteFromDeck(deckName)
	t.Equal(nil, err)
	t.Equal([]anki2.Note{domainNote}, noteList)
}

func (t *AnkiSuite) Test_update_note_and_audio() {
	var err error
	domainNote, notesInfo := t.getInfrastructureAndDomainNote()
	notesInfo.NoteId = 456
	t.mockAnkier.EXPECT().EditNoteById(notesInfo, []ankiconnect.Audio{{
		URL:      "url audio link",
		Data:     "base64 audio",
		Path:     "audio absolute path",
		Filename: "fileName",
		SkipHash: "",
		Fields:   []string{"Meaning"},
	}}, nil, nil)
	err = t.anki.UpdateNoteById(456, domainNote, []anki2.Audio{
		{
			URL:      "url audio link",
			Data:     "base64 audio",
			Path:     "audio absolute path",
			Filename: "fileName",
			SkipHash: "",
			Fields:   []string{"Meaning"},
		},
	})
	t.NoError(err)
}

func (t *AnkiSuite) getInfrastructureAndDomainNote() (anki2.Note, ankiconnect.ResultNotesInfo) {
	noteId := int64(123)
	modelName := "model"
	fieldData := map[string]struct {
		Value string
		Order int64
	}{
		"Meaning": {
			"Meaning Value",
			0,
		},
	}
	tags := []string{"tag1", "tag2"}
	exampleNote := anki2.Note{
		Id:        noteId,
		ModelName: modelName,
		Fields:    map[string]anki2.FieldData{},
		Tags:      tags,
	}
	exampleResultNotesInfo := ankiconnect.ResultNotesInfo{
		NoteId:    noteId,
		ModelName: modelName,
		Fields:    map[string]ankiconnect.FieldData{},
		Tags:      tags,
	}
	for key, data := range fieldData {
		exampleNote.Fields[key] = anki2.FieldData{
			data.Value, data.Order,
		}
		exampleResultNotesInfo.Fields[key] = ankiconnect.FieldData{
			data.Value, data.Order,
		}
	}
	return exampleNote, exampleResultNotesInfo
}
