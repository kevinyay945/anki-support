package application

import (
	"anki-support/domain"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type AnkiRepositorySuite struct {
	suite.Suite
	mockCtrl       *gomock.Controller
	ankiRepository *AnkiRepository
	mockAnkier     *domain.MockAnkier
}

func TestSuiteInitAnkiRepository(t *testing.T) {
	suite.Run(t, new(AnkiRepositorySuite))
}

func (t *AnkiRepositorySuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.mockAnkier = domain.NewMockAnkier(t.mockCtrl)
	t.ankiRepository = &AnkiRepository{
		ankier: t.mockAnkier,
	}
}

func (t *AnkiRepositorySuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *AnkiRepositorySuite) Test_Get_anki_note_by_deck_name() {
	testAnkiNoteList := []domain.AnkiNote{
		{
			Id:        0,
			ModelName: "",
			Fields:    nil,
			Tags:      nil,
		},
	}
	t.mockAnkier.EXPECT().GetNoteListByDeckName("Test Deck Name").Return(
		testAnkiNoteList,
		nil,
	)
	noteList, err := t.ankiRepository.GetAllNotesByDeckName("Test Deck Name")
	t.NoError(err)
	t.Equal(noteList, testAnkiNoteList)
}

func (t *AnkiRepositorySuite) Test_Get_anki_todo_note_by_deck_name() {
	testAnkiNoteList := []domain.AnkiNote{
		{
			Id:        0,
			ModelName: "",
			Fields:    nil,
			Tags:      nil,
		},
	}
	t.mockAnkier.EXPECT().GetTodoNoteFromDeck("Test Deck Name").Return(testAnkiNoteList, nil)
	noteList, err := t.ankiRepository.GetAllTodoNotesByDeckName("Test Deck Name")
	t.NoError(err)
	t.Equal(noteList, testAnkiNoteList)
}
