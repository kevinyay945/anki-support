package domain

import (
	"anki-support/infrastructure/anki"
	"github.com/atselvan/ankiconnect"
	"regexp"
)

type (
	Note struct {
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

type Anki struct {
	anki anki.Ankier
}

func NewAnki(anki anki.Ankier) Ankier {
	return &Anki{anki: anki}
}

type Ankier interface {
	GetNoteListByDeckName(deckName string) (output []Note, err error)
	GetNoteById(noteId int64) (output Note, err error)
	GetTodoNoteFromDeck(deckName string) (output []Note, err error)
	UpdateNoteById(noteId int64, note Note, audioList []Audio) (err error)
}

func (a *Anki) UpdateNoteById(noteId int64, note Note, audioList []Audio) (err error) {
	updateNote := note.ankiconnectNoteInfo()
	updateNote.NoteId = noteId
	var updateAudioList []ankiconnect.Audio
	for _, audio := range audioList {
		updateAudioList = append(updateAudioList, audio.ankiconnectAudio())
	}
	err = a.anki.EditNoteById(updateNote, updateAudioList, nil, nil)
	return
}

func (audio Audio) ankiconnectAudio() ankiconnect.Audio {
	elems := ankiconnect.Audio{
		URL:      audio.URL,
		Data:     audio.Data,
		Path:     audio.Path,
		Filename: audio.Filename,
		SkipHash: audio.SkipHash,
		Fields:   audio.Fields,
	}
	return elems
}

func (n Note) ankiconnectNoteInfo() ankiconnect.ResultNotesInfo {
	updateNote := ankiconnect.ResultNotesInfo{
		NoteId:    n.Id,
		ModelName: n.ModelName,
		Fields:    map[string]ankiconnect.FieldData{},
		Tags:      n.Tags,
	}
	for key, data := range n.Fields {
		updateNote.Fields[key] = ankiconnect.FieldData{
			Value: data.Value, Order: data.Order,
		}
	}
	return updateNote
}

func (a *Anki) GetTodoNoteFromDeck(deckName string) (output []Note, err error) {
	noteList, err := a.anki.GetTodoNoteFromDeck(deckName)
	for _, note := range noteList {
		output = append(output, Note{}.FromResultNotesInfo(note))
	}
	return
}

func (a *Anki) GetNoteById(noteId int64) (output Note, err error) {
	note, err := a.anki.GetNoteById(noteId)
	return Note{}.FromResultNotesInfo(note), err
}

func (a *Anki) GetNoteListByDeckName(s string) (output []Note, err error) {
	noteList, err := a.anki.GetAllNoteFromDeck(s)
	for _, note := range noteList {
		output = append(output, Note{}.FromResultNotesInfo(note))
	}
	return
}
func (n Note) FromResultNotesInfo(note ankiconnect.ResultNotesInfo) (output Note) {
	output.ModelName = note.ModelName
	output.Tags = note.Tags
	output.Fields = map[string]FieldData{}
	for key, data := range note.Fields {
		output.Fields[key] = FieldData{
			Value: data.Value,
			Order: data.Order,
		}
	}
	output.Id = note.NoteId
	return
}

func (n Note) HasSound(s string) bool {
	value := n.Fields[s].Value
	pattern := `\[sound:([^\]]+)\]`
	re := regexp.MustCompile(pattern)
	matches := re.FindAllStringSubmatch(value, -1)
	return matches != nil
}

func (n Note) HasValue(s string) bool {
	value := n.Fields[s].Value
	return value != ""
}

func (n Note) GetValue(s string) string {
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
