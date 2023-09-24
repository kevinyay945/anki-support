package domain

import (
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
	note := AnkiNote{
		ModelName: "Japanese (recognition&recall)",
	}
	generator := OperatorGenerate{}
	normalOperator, _ := generator.GetByNote(note)
	_, ok := normalOperator.(*NormalOperator)
	t.True(ok, "type is not *NormalOperator")

	note = AnkiNote{
		ModelName: "Japanese (recognition&recall) 動詞篇",
	}
	generator = OperatorGenerate{}
	verbOperator, _ := generator.GetByNote(note)
	_, ok = verbOperator.(*VerbOperator)
	t.True(ok, "type is not *VerbOperator")

	note = AnkiNote{
		ModelName: "Japanese (recognition&recall) 形容詞",
	}
	generator = OperatorGenerate{}
	adjOperator, _ := generator.GetByNote(note)
	_, ok = adjOperator.(*AdjOperator)
	t.True(ok, "type is not *AdjOperator")
}
