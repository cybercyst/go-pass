package commands

import (
	"github.com/urfave/cli/v3"
)

var Remove = &cli.Command{
	Name:   "rm",
	Action: stubAction,
}
