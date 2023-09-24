package anki

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

func (t *AnkiOperationSuite) Test_get_correct_modal_type() {
	// get all note from deck

	// check correct type and return correct operator
}
