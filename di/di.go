package di

import (
	"anki-support/interfaces/cmd"
	_ "github.com/google/subcommands"
	_ "github.com/google/wire"
)

type DI struct {
	RunCmd *cmd.RunCmd
}

func NewDI(runCmd *cmd.RunCmd) *DI {
	return &DI{RunCmd: runCmd}
}
