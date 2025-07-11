package commands

import (
	"github.com/urfave/cli/v3"
)

var Git = &cli.Command{
	Name:   "git",
	Action: stubAction,
}
