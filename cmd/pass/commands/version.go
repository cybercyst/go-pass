package commands

import (
	"github.com/urfave/cli/v3"
)

var Version = &cli.Command{
	Name:   "version",
	Action: stubAction,
}
