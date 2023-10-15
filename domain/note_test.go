package domain

import (
	"github.com/stretchr/testify/suite"
	"go.uber.org/mock/gomock"
	"testing"
)

type NoteSuite struct {
	suite.Suite
	mockCtrl *gomock.Controller
}

func TestSuiteInitNote(t *testing.T) {
	suite.Run(t, new(NoteSuite))
}

func (t *NoteSuite) SetupTest() {
	t.mockCtrl = gomock.NewController(t.Suite.T())
}

func (t *NoteSuite) TearDownTest() {
	defer t.mockCtrl.Finish()
}

func (t *NoteSuite) Test_anki_note_column_has_sound() {
	note := AnkiNote{
		Fields: map[string]AnkiFieldData{
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

func (t *NoteSuite) Test_anki_note_column_has_value() {
	n := AnkiNote{
		Fields: map[string]AnkiFieldData{
			"hasValue": {
				Value: "there are some value",
			},
			"noValue": {
				Value: "",
			},
		},
	}
	t.Equal(true, n.HasValue("hasValue"))
	t.Equal(false, n.HasValue("noValue"))
}

func (t *NoteSuite) Test_anki_note_column_get_column_and_get_value_only() {
	note := AnkiNote{
		Fields: map[string]AnkiFieldData{
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
