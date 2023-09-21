package domain

import (
	"anki-support/infrastructure/anki"
	"github.com/atselvan/ankiconnect"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type AnkiSuite struct {
	suite.Suite
	mockCtrl   *gomock.Controller
	anki       Ankier
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
	t.Equal([]Note{exampleNote}, note)
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
	t.Equal([]Note{domainNote}, noteList)
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
	err = t.anki.UpdateNoteById(456, domainNote, []Audio{
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

func (t *AnkiSuite) Test_anki_note_column_has_sound() {
	note := Note{
		Fields: map[string]FieldData{
			"hasSoundField": {
				Value: "[sound:test.mp3]",
			},
			"noSoundField": {
				Value: "no sound",
			},
		},
	}
	t.Equal(true, note.HasSound("hasSoundField"))
	t.Equal(false, note.HasSound("noSoundField"))
}

func (t *AnkiSuite) Test_anki_note_column_has_value() {
	note := Note{
		Fields: map[string]FieldData{
			"hasValue": {
				Value: "there are some value",
			},
			"noValue": {
				Value: "",
			},
		},
	}
	t.Equal(true, note.HasValue("hasValue"))
	t.Equal(false, note.HasValue("noValue"))
}

func (t *AnkiSuite) Test_anki_note_column_get_column_and_get_value_only() {
	note := Note{
		Fields: map[string]FieldData{
			"target": {
				Value: "エアコン<br><!-- user_accent_start --><br><hr><br><svg class=\\\"pitch\\\" width=\\\"172px\\\" height=\\\"75px\\\" viewBox=\\\"0 0 172 75\\\"><text x=\\\"5\\\" y=\\\"67.5\\\" style=\\\"font-size:20px;font-family:sans-serif;fill:#000;\\\">エ</text><text x=\\\"40\\\" y=\\\"67.5\\\" style=\\\"font-size:20px;font-family:sans-serif;fill:#000;\\\">ア</text><text x=\\\"75\\\" y=\\\"67.5\\\" style=\\\"font-size:20px;font-family:sans-serif;fill:#000;\\\">コ</text><text x=\\\"110\\\" y=\\\"67.5\\\" style=\\\"font-size:20px;font-family:sans-serif;fill:#000;\\\">ン</text><path d=\\\"m 16,30 35,-25\\\" style=\\\"fill:none;stroke:#000;stroke-width:1.5;\\\"></path><path d=\\\"m 51,5 35,0\\\" style=\\\"fill:none;stroke:#000;stroke-width:1.5;\\\"></path><path d=\\\"m 86,5 35,0\\\" style=\\\"fill:none;stroke:#000;stroke-width:1.5;\\\"></path><path d=\\\"m 121,5 35,0\\\" style=\\\"fill:none;stroke:#000;stroke-width:1.5;\\\"></path><circle r=\\\"5\\\" cx=\\\"16\\\" cy=\\\"30\\\" style=\\\"opacity:1;fill:#000;\\\"></circle><circle r=\\\"5\\\" cx=\\\"51\\\" cy=\\\"5\\\" style=\\\"opacity:1;fill:#000;\\\"></circle><circle r=\\\"5\\\" cx=\\\"86\\\" cy=\\\"5\\\" style=\\\"opacity:1;fill:#000;\\\"></circle><circle r=\\\"5\\\" cx=\\\"121\\\" cy=\\\"5\\\" style=\\\"opacity:1;fill:#000;\\\"></circle><circle r=\\\"5\\\" cx=\\\"156\\\" cy=\\\"5\\\" style=\\\"opacity:1;fill:#000;\\\"></circle><circle r=\\\"3.25\\\" cx=\\\"156\\\" cy=\\\"5\\\" style=\\\"opacity:1;fill:#fff;\\\"></circle></svg><!-- user_accent_end -->",
			},
			"normal": {
				Value: "時計<hr><br />",
			},
		},
	}
	t.Equal("エアコン", note.GetValue("target"))
	t.Equal("時計", note.GetValue("normal"))
}

func (t *AnkiSuite) getInfrastructureAndDomainNote() (Note, ankiconnect.ResultNotesInfo) {
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
	exampleNote := Note{
		Id:        noteId,
		ModelName: modelName,
		Fields:    map[string]FieldData{},
		Tags:      tags,
	}
	exampleResultNotesInfo := ankiconnect.ResultNotesInfo{
		NoteId:    noteId,
		ModelName: modelName,
		Fields:    map[string]ankiconnect.FieldData{},
		Tags:      tags,
	}
	for key, data := range fieldData {
		exampleNote.Fields[key] = FieldData{
			data.Value, data.Order,
		}
		exampleResultNotesInfo.Fields[key] = ankiconnect.FieldData{
			data.Value, data.Order,
		}
	}
	return exampleNote, exampleResultNotesInfo
}
