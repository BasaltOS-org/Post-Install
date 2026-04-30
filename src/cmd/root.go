package cmd

import (
	"github.com/urfave/cli/v3"
)

func InitCmds() (cmds []*cli.Command) {
	cmds = []*cli.Command{
		cmdPackage(),
	}
	return cmds
}
