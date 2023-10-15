//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.
package di

import (
	"anki-support/application"
	"anki-support/infrastructure"
	"anki-support/interfaces/cmd"
	"anki-support/lib/anki"
	"anki-support/lib/gcp"
	"anki-support/lib/log"
	"anki-support/lib/openai"
	"github.com/google/wire"
)

// InitializeAuthCmd creates an Auth Init Struct. It will error if the Event is staffed with
// a grumpy greeter.
func InitializeDICmd() *DI {
	wire.Build(
		log.NewLogger,
		gcp.NewGCP,
		openai.NewClient,
		anki.NewClient,
		infrastructure.NewAnki,
		infrastructure.NewGPT,
		infrastructure.NewTextToSpeech,
		application.NewAnkiOperatorFactory,
		application.NewAnkiRepository,
		cmd.NewRunCmd,
		NewDI,
	)
	return nil
}
