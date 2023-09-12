package anki

import (
	"github.com/atselvan/ankiconnect"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"os"
	"testing"
)

type AnkiSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
	client   *Client
}

func TestSuiteInitAnki(t *testing.T) {
	if os.Getenv("RUN_INFRASTRUCTURE") == "true" {
		t.Skip("Skipping testing in production")
	}
}

func (t *AnkiSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	client := NewClient()
	t.client = client
}

func (t *AnkiSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *AnkiSuite) Test_ping_anki_is_alive() {
	err := t.client.Ping()
	t.NoError(err)
}

func (t *AnkiSuite) Test_get_deck_list() {
	deck, err := t.client.GetAllDeck()
	t.NoError(err)
	t.Contains(deck, "製作中日語卡片")
}

func (t *AnkiSuite) Test_get_all_note_from_deck_name() {
	deckName := "製作中日語卡片"
	note, err := t.client.GetAllNoteFromDeck(deckName)
	t.NoError(err)
	t.Contains(note, ankiconnect.Note{})
}

func (t *AnkiSuite) Test_get_todo_note_from_deck_name() {
	deckName := "製作中日語卡片"
	note, err := t.client.GetTodoNoteFromDeck(deckName)
	t.NoError(err)
	t.Len(note, 2)
	t.Equal(12345, note[0].NoteId)
	t.Equal("aaa", note[0].ModelName)
}

func (t *AnkiSuite) Test_get_note_by_id() {
	noteId := int64(1694305287189)
	_, err := t.client.GetNoteById(noteId)
	t.NoError(err)
}

func (t *AnkiSuite) Test_edit_note_and_add_audio() {
	noteId := int64(1694305287189)
	note, err := t.client.GetNoteById(noteId)
	t.NoError(err)
	outputPath := "/Users/kevin/Developer/side-project/anki-support/asset/test/私の机は木製です。.mp3"
	var audioList = []ankiconnect.Audio{
		{
			URL:      "",
			Data:     "",
			Path:     outputPath,
			Filename: "私の机は木製です。.mp3",
			SkipHash: "",
			Fields:   []string{"JapaneseSound"},
		},
	}
	err = t.client.EditNoteById(note, audioList, nil, nil)
	t.NoError(err)
}

func (t *AnkiSuite) Test_delete_tag_from_note() {
	noteId := int64(1694305287189)
	err := t.client.DeleteTagFromNote(noteId, "anki-helper-vocabulary-done")
	t.NoError(err)
}

func (t *AnkiSuite) Test_add_tag_from_note() {
	noteId := int64(1694305287189)
	err := t.client.AddTagFromNote(noteId, "anki-helper-vocabulary-todo")
	t.NoError(err)
}
