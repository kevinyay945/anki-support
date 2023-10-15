package application

//go:generate mockgen -destination=anki_operator.mock.go -typed=true -package=application -self_package=anki-support/application . AnkiOperator
type AnkiOperator interface {
	Do() error
}
