package cmd

import (
	"anki-support/application"
	"anki-support/domain"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type RunSuite struct {
	suite.Suite
	mockCtrl                  *gomock.Controller
	runCmd                    *RunCmd
	mockAnkiOperator          *application.MockAnkiOperator
	mockAnkiOperatorFactorier *application.MockAnkiOperatorFactorier
	mockAnkiRepositorier      *application.MockAnkiRepositorier
}

func TestSuiteInitRun(t *testing.T) {
	suite.Run(t, new(RunSuite))
}

func (t *RunSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
	t.mockAnkiOperator = application.NewMockAnkiOperator(t.mockCtrl)
	t.mockAnkiOperatorFactorier = application.NewMockAnkiOperatorFactorier(t.mockCtrl)
	t.mockAnkiRepositorier = application.NewMockAnkiRepositorier(t.mockCtrl)
	t.runCmd = &RunCmd{
		ankiOperatorFactory: t.mockAnkiOperatorFactorier,
		ankiRepository:      t.mockAnkiRepositorier,
	}
}

func (t *RunSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *RunSuite) Test_run_for_specific_deck() {
	todoNote := domain.AnkiNote{
		Id:        0,
		ModelName: "",
		Fields: map[string]domain.AnkiFieldData{
			"Expression": domain.AnkiFieldData{
				Value: "todo value",
				Order: 0,
			},
		},
		Tags: nil,
	}
	finishNote := domain.AnkiNote{
		Id:        1,
		ModelName: "",
		Fields: map[string]domain.AnkiFieldData{
			"Expression": domain.AnkiFieldData{
				Value: "finish value",
				Order: 0,
			},
		},
		Tags: nil,
	}
	t.mockAnkiRepositorier.EXPECT().
		GetAllTodoNotesByDeckName("test_deck_name").
		Return([]domain.AnkiNote{
			todoNote,
		}, nil)
	t.mockAnkiRepositorier.EXPECT().
		GetAllNotesByDeckName("test_deck_name").
		Return([]domain.AnkiNote{
			todoNote,
			finishNote,
		}, nil)
	t.mockAnkiOperatorFactorier.EXPECT().CreateByNote(
		todoNote,
		[]string{
			todoNote.Fields["Expression"].Value,
			finishNote.Fields["Expression"].Value,
		},
	).Return(t.mockAnkiOperator, nil)
	t.mockAnkiOperator.EXPECT().Do().Times(1).Return(nil)
	err := t.runCmd.RunForSpecificDeck("test_deck_name")
	t.NoError(err)
}
