package infrastructure

import (
	"anki-support/domain"
	"anki-support/lib/anki"
	"github.com/atselvan/ankiconnect"
)

type Anki struct {
	anki anki.Ankier
}

func (a *Anki) AddNoteTagFromNoteId(NoteId int64, tagName string) (err error) {
	err = a.anki.AddTagFromNote(NoteId, tagName)
	return
}

func (a *Anki) DeleteNoteTagFromNoteId(NoteId int64, tagName string) (err error) {
	err = a.anki.DeleteTagFromNote(NoteId, tagName)
	return
}

func NewAnki(anki anki.Ankier) domain.Ankier {
	return &Anki{anki: anki}
}

func (a *Anki) UpdateNoteById(noteId int64, note domain.AnkiNote, audioList []domain.Audio) (err error) {
	updateNote := DomainNoteToAnkiconnectNoteInfo(note)
	updateNote.NoteId = noteId
	var updateAudioList []ankiconnect.Audio
	for _, audio := range audioList {
		updateAudioList = append(updateAudioList, DomainAudioToAnkiconnectAudio(audio))
	}
	err = a.anki.EditNoteById(updateNote, updateAudioList, nil, nil)
	return
}

func DomainAudioToAnkiconnectAudio(audio domain.Audio) ankiconnect.Audio {
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

func DomainNoteToAnkiconnectNoteInfo(n domain.AnkiNote) ankiconnect.ResultNotesInfo {
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

func (a *Anki) GetTodoNoteFromDeck(deckName string) (output []domain.AnkiNote, err error) {
	noteList, err := a.anki.GetNoteFromDeckByTagName(deckName, domain.AnkiTodoTagName)
	for _, note := range noteList {
		output = append(output, GetNoteFromResultNotesInfo(note))
	}
	return
}

func (a *Anki) GetNoteById(noteId int64) (output domain.AnkiNote, err error) {
	note, err := a.anki.GetNoteById(noteId)
	return GetNoteFromResultNotesInfo(note), err
}

func (a *Anki) GetNoteListByDeckName(s string) (output []domain.AnkiNote, err error) {
	noteList, err := a.anki.GetAllNoteFromDeck(s)
	for _, note := range noteList {
		output = append(output, GetNoteFromResultNotesInfo(note))
	}
	return
}

func GetNoteFromResultNotesInfo(note ankiconnect.ResultNotesInfo) (output domain.AnkiNote) {
	output.ModelName = note.ModelName
	output.Tags = note.Tags
	output.Fields = map[string]domain.FieldData{}
	for key, data := range note.Fields {
		output.Fields[key] = domain.FieldData{
			Value: data.Value,
			Order: data.Order,
		}
	}
	output.Id = note.NoteId
	return
}
