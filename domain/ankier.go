package domain

import "regexp"

//go:generate mockgen -destination=ankier.mock.go -typed=true -package=domain -self_package=anki-support/domain . Ankier
type Ankier interface {
	GetNoteListByDeckName(deckName string) (output []AnkiNote, err error)
	GetNoteById(noteId int64) (output AnkiNote, err error)
	GetTodoNoteFromDeck(deckName string) (output []AnkiNote, err error)
	AddNoteTagFromNoteId(NoteId int64, tagName string) (err error)
	DeleteNoteTagFromNoteId(NoteId int64, tagName string) (err error)
	// UpdateNoteById can't update tag at the same time
	UpdateNoteById(noteId int64, note AnkiNote, audioList []Audio) (err error)
}

type (
	AnkiNote struct {
		Id        int64
		ModelName string
		Fields    map[string]FieldData
		Tags      []string
	}
	FieldData struct {
		Value string
		Order int64
	}

	// Audio can be used to add a audio file to a Anki Note.
	Audio struct {
		URL      string   `json:"url,omitempty"`
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Video can be used to add a video file to a Anki Note.
	Video struct {
		URL      string   `json:"url,omitempty"`
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}

	// Picture can be used to add a picture to an Anki Note.
	Picture struct {
		URL      string   `json:"url,omitempty"`
		Data     string   `json:"data,omitempty"`
		Path     string   `json:"path,omitempty"`
		Filename string   `json:"filename,omitempty"`
		SkipHash string   `json:"skipHash,omitempty"`
		Fields   []string `json:"fields,omitempty"`
	}
)

const AnkiDoneTagName = "anki-helper-vocabulary-done"

const AnkiTodoTagName = "anki-helper-vocabulary-todo"

func (n AnkiNote) HasSound(s string) bool {
	value := n.Fields[s].Value
	pattern := `\[sound:([^\]]+)\]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(value, -1)
	return matches != nil
}

func (n AnkiNote) HasValue(s string) bool {
	value := n.Fields[s].Value
	return value != ""
}

func (n AnkiNote) GetValue(s string) string {
	value := n.Fields[s].Value

	// 使用正規表達式來刪除 <!-- user_accent_start --> 和 <!-- user_accent_end --> 之間的內容
	pattern := `<!-- user_accent_start -->(.*?)<!-- user_accent_end -->`
	re := regexp.MustCompile(pattern)
	result := re.ReplaceAllString(value, "")

	// 刪除HTML標籤
	reHTML := regexp.MustCompile(`<[^>]+>`)
	result = reHTML.ReplaceAllString(result, "")

	return result
}
