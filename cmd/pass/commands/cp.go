package commands

import (
	"github.com/urfave/cli/v3"
)

var Copy = &cli.Command{
	Name:   "cp",
	Action: stubAction,
}
