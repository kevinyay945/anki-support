package anki

import (
	"fmt"
	"github.com/atselvan/ankiconnect"
	"github.com/imroc/req/v3"
	"github.com/privatesquare/bkst-go-utils/utils/errors"
	"github.com/sirupsen/logrus"
	"os"
	"path/filepath"
)

type Client struct {
	ankiConnect   *ankiconnect.Client
	ankiMediaPath string
	httpClient    *req.Client
	log           *logrus.Entry
}

func NewClient(log *logrus.Logger) *Client {
	field := log.WithField("infrastructure", "anki")
	c := &Client{log: field}
	c.ankiConnect = ankiconnect.NewClient()
	c.httpClient = req.C()
	c.httpClient.SetBaseURL("http://127.0.0.1:8765")
	if path, err := c.GetMediaFolderPath(); err != nil {
		c.log.Warnf("Fail To Get Media Folder: %s", err.Error())
	} else {
		c.ankiMediaPath = path
	}
	return c
}

func (c *Client) Ping() (err error) {
	restErr := c.ankiConnect.Ping()
	return NewAnkiError(restErr)
}

func (c *Client) GetAllDeck() ([]string, error) {
	all, err := c.ankiConnect.Decks.GetAll()
	return *all, NewAnkiError(err)
}

func (c *Client) GetAllNoteFromDeck(name string) ([]ankiconnect.ResultNotesInfo, error) {
	get, err := c.ankiConnect.Notes.Get(fmt.Sprintf("deck:%s", name))
	return *get, NewAnkiError(err)
}

func (c *Client) GetTodoNoteFromDeck(deckName string) ([]ankiconnect.ResultNotesInfo, error) {
	todoTag := "anki-helper-vocabulary-todo"
	get, err := c.ankiConnect.Notes.Get(fmt.Sprintf("tag:%s deck:%s", todoTag, deckName))
	return *get, NewAnkiError(err)
}

// EditNoteById should be careful that you can't edit tag, and you can't edit when you open this card on anki gui
func (c *Client) EditNoteById(
	note ankiconnect.ResultNotesInfo,
	audioList []ankiconnect.Audio,
	videoList []ankiconnect.Video,
	pictureList []ankiconnect.Picture,
) error {
	updateFields := map[string]string{}
	for s, data := range note.Fields {
		updateFields[s] = data.Value
	}
	var oldAudioList []ankiconnect.Audio
	for _, audio := range audioList {
		if audio.Path == "" {
			oldAudioList = append(oldAudioList, audio)
			continue
		}
		input, err := os.ReadFile(audio.Path)
		if err != nil {
			c.log.Warnf("Fail to read audio file: %s", err.Error())
			oldAudioList = append(oldAudioList, audio)
			continue
		}

		ankiAudioFile := fmt.Sprintf("anki-support-%s", audio.Filename)
		destinationFile := filepath.Join(c.ankiMediaPath, ankiAudioFile)
		err = os.WriteFile(destinationFile, input, 0644)
		if err != nil {
			c.log.Warnln("Error creating", destinationFile)
			c.log.Warnln(err.Error())
			oldAudioList = append(oldAudioList, audio)
			continue
		}
		for _, field := range audio.Fields {
			updateFields[field] = fmt.Sprintf("[sound:%s]", ankiAudioFile)
		}
	}
	err := c.ankiConnect.Notes.Update(ankiconnect.UpdateNote{
		Id:      note.NoteId,
		Fields:  updateFields,
		Audio:   oldAudioList,
		Video:   videoList,
		Picture: pictureList,
	})
	return NewAnkiError(err)
}

func (c *Client) GetNoteById(id int64) (ankiconnect.ResultNotesInfo, error) {
	result := struct {
		Result []struct {
			NoteId int64    `json:"noteId"`
			Tags   []string `json:"tags"`
			Fields map[string]struct {
				Value string `json:"value"`
				Order int64  `json:"order"`
			} `json:"fields"`
			ModelName string `json:"modelName"`
		} `json:"result"`
		Err string `json:"error"`
	}{}
	errResult := struct {
		Err string `json:"error"`
	}{}
	post, err := c.httpClient.R().SetBodyJsonString(fmt.Sprintf(`
	{
		"action": "notesInfo",
		"version": 6,
		"params": {
			"notes": [%d]
		}
	}
	`, id)).
		SetSuccessResult(&result).
		SetErrorResult(&errResult).
		Post("")
	if err != nil {
		return ankiconnect.ResultNotesInfo{}, err
	}
	if post.IsErrorState() {
		return ankiconnect.ResultNotesInfo{}, fmt.Errorf("fail to get note by id")
	}
	if result.Err != "" {
		return ankiconnect.ResultNotesInfo{}, fmt.Errorf("fail to get note by id")
	}
	firstNote := result.Result[0]
	output := ankiconnect.ResultNotesInfo{}
	output.NoteId = firstNote.NoteId
	output.ModelName = firstNote.ModelName
	output.Tags = firstNote.Tags
	output.Fields = map[string]ankiconnect.FieldData{}
	for key, value := range firstNote.Fields {
		output.Fields[key] = ankiconnect.FieldData{
			Value: value.Value,
			Order: value.Order,
		}
	}
	return output, nil
}

func (c *Client) DeleteTagFromNote(noteId int64, tag string) error {
	result := struct {
		Error string `json:"error"`
	}{}
	post, err := c.httpClient.R().SetBodyJsonString(fmt.Sprintf(`
	{
		"action": "removeTags",
		"version": 6,
		"params": {
			"notes": [%d],
			"tags": "%s"
		}
	}
	`, noteId, tag)).
		SetSuccessResult(&result).
		Post("")
	if err != nil {
		return fmt.Errorf("fail to delete note tag")
	}
	if post.IsErrorState() {
		return fmt.Errorf("fail to delete note tag")
	}
	if result.Error != "" {
		return fmt.Errorf("fail to delete note tag")
	}
	return nil
}

func (c *Client) AddTagFromNote(id int64, tag string) error {
	result := struct {
		Error string `json:"error"`
	}{}
	post, err := c.httpClient.R().SetBodyJsonString(fmt.Sprintf(`
	{
		"action": "addTags",
		"version": 6,
		"params": {
			"notes": [%d],
			"tags": "%s"
		}
	}
	`, id, tag)).
		SetSuccessResult(&result).
		Post("")
	if err != nil {
		return fmt.Errorf("fail to add note tag")
	}
	if post.IsErrorState() {
		return fmt.Errorf("fail to add note tag")
	}
	if result.Error != "" {
		return fmt.Errorf("fail to add note tag")
	}
	return nil
}

func (c *Client) GetMediaFolderPath() (string, error) {
	result := struct {
		Result string `json:"result"`
		Error  string `json:"error"`
	}{}
	post, err := c.httpClient.R().SetBodyJsonString(`
	{
		"action": "getMediaDirPath",
		"version": 6
	}
	`).SetSuccessResult(&result).Post("")
	if err != nil {
		return "", fmt.Errorf("fail to add note tag")
	}
	if post.IsErrorState() {
		return "", fmt.Errorf("fail to add note tag")
	}
	if result.Error != "" {
		return "", fmt.Errorf("fail to add note tag")
	}
	return result.Result, nil
}

func NewAnkiError(err *errors.RestErr) error {
	if err == nil {
		return nil
	}
	//Message    string `json:"message"`
	//StatusCode int    `json:"status"`
	//Error      string `json:"error"`
	return fmt.Errorf("status: %d, message: %s, error: %s", err.StatusCode, err.Message, err.Error)
}

type Clienter interface {
	Ping() (err error)
	GetAllDeck() ([]string, error)
	GetAllNoteFromDeck(name string) ([]ankiconnect.ResultNotesInfo, error)
	GetTodoNoteFromDeck(deckName string) ([]ankiconnect.ResultNotesInfo, error)
	EditNoteById(
		note ankiconnect.ResultNotesInfo,
		audioList []ankiconnect.Audio,
		videoList []ankiconnect.Video,
		pictureList []ankiconnect.Picture,
	) error
	GetNoteById(id int64) (ankiconnect.ResultNotesInfo, error)
	DeleteTagFromNote(noteId int64, tag string) error
	AddTagFromNote(id int64, tag string) error
}
