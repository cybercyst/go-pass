package commands

import (
	"github.com/urfave/cli/v3"
)

var Move = &cli.Command{
	Name:   "mv",
	Action: stubAction,
}
