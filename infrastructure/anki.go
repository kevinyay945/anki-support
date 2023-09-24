package infrastructure

import (
	anki2 "anki-support/domain/anki"
	"anki-support/lib/anki"
	"github.com/atselvan/ankiconnect"
)

type Anki struct {
	anki anki.Ankier
}

func NewAnki(anki anki.Ankier) anki2.Ankier {
	return &Anki{anki: anki}
}

func (a *Anki) UpdateNoteById(noteId int64, note anki2.Note, audioList []anki2.Audio) (err error) {
	updateNote := DomainNoteToAnkiconnectNoteInfo(note)
	updateNote.NoteId = noteId
	var updateAudioList []ankiconnect.Audio
	for _, audio := range audioList {
		updateAudioList = append(updateAudioList, DomainAudioToAnkiconnectAudio(audio))
	}
	err = a.anki.EditNoteById(updateNote, updateAudioList, nil, nil)
	return
}

func DomainAudioToAnkiconnectAudio(audio anki2.Audio) ankiconnect.Audio {
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

func DomainNoteToAnkiconnectNoteInfo(n anki2.Note) ankiconnect.ResultNotesInfo {
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

func (a *Anki) GetTodoNoteFromDeck(deckName string) (output []anki2.Note, err error) {
	noteList, err := a.anki.GetTodoNoteFromDeck(deckName)
	for _, note := range noteList {
		output = append(output, GetNoteFromResultNotesInfo(note))
	}
	return
}

func (a *Anki) GetNoteById(noteId int64) (output anki2.Note, err error) {
	note, err := a.anki.GetNoteById(noteId)
	return GetNoteFromResultNotesInfo(note), err
}

func (a *Anki) GetNoteListByDeckName(s string) (output []anki2.Note, err error) {
	noteList, err := a.anki.GetAllNoteFromDeck(s)
	for _, note := range noteList {
		output = append(output, GetNoteFromResultNotesInfo(note))
	}
	return
}

func GetNoteFromResultNotesInfo(note ankiconnect.ResultNotesInfo) (output anki2.Note) {
	output.ModelName = note.ModelName
	output.Tags = note.Tags
	output.Fields = map[string]anki2.FieldData{}
	for key, data := range note.Fields {
		output.Fields[key] = anki2.FieldData{
			Value: data.Value,
			Order: data.Order,
		}
	}
	output.Id = note.NoteId
	return
}
