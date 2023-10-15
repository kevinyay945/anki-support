package application

import "anki-support/domain"

type AnkiRepository struct {
	ankier domain.Ankier
}

func (r *AnkiRepository) GetAllNotesByDeckName(deckName string) ([]domain.AnkiNote, error) {
	output, err := r.ankier.GetNoteListByDeckName(deckName)
	return output, err
}

func (r *AnkiRepository) GetAllTodoNotesByDeckName(deckName string) ([]domain.AnkiNote, error) {
	output, err := r.ankier.GetTodoNoteFromDeck(deckName)
	return output, err
}

//go:generate mockgen -destination=anki_operator_repository.mock.go -typed=true -package=application -self_package=anki-support/application . AnkiRepositorier
type AnkiRepositorier interface {
	GetAllNotesByDeckName(deckName string) ([]domain.AnkiNote, error)
	GetAllTodoNotesByDeckName(deckName string) ([]domain.AnkiNote, error)
}
