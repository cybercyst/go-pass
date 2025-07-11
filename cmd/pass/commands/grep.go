package commands

import (
	"github.com/urfave/cli/v3"
)

var Grep = &cli.Command{
	Name:   "grep",
	Action: stubAction,
}
