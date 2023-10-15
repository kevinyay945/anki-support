package application

import (
	"anki-support/domain"
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type AnkiOperationSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestSuiteInitAnkiOperation(t *testing.T) {
	suite.Run(t, new(AnkiOperationSuite))
}

func (t *AnkiOperationSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
}

func (t *AnkiOperationSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *AnkiOperationSuite) Test_get_correct_modal_type_operator() {
	note := domain.AnkiNote{
		ModelName: "Japanese (recognition&recall)",
	}
	generator := AnkiOperatorFactory{}
	normalOperator, _ := generator.CreateByNote(note, nil)
	_, ok := normalOperator.(*AnkiNormalJapaneseOperator)
	t.True(ok, "type is not *AnkiNormalJapaneseOperator")

	note = domain.AnkiNote{
		ModelName: "Japanese (recognition&recall) 動詞篇",
	}
	generator = AnkiOperatorFactory{}
	verbOperator, _ := generator.CreateByNote(note, nil)
	_, ok = verbOperator.(*AnkiVerbOperator)
	t.True(ok, "type is not *AnkiVerbOperator")

	note = domain.AnkiNote{
		ModelName: "Japanese (recognition&recall) 形容詞",
	}
	generator = AnkiOperatorFactory{}
	adjOperator, _ := generator.CreateByNote(note, nil)
	_, ok = adjOperator.(*AnkiAdjOperator)
	t.True(ok, "type is not *AnkiAdjOperator")
}
