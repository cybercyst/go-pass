package commands

import (
	"github.com/urfave/cli/v3"
)

var Generate = &cli.Command{
	Name:   "generate",
	Action: stubAction,
}
