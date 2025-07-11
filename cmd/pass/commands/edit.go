package commands

import (
	"github.com/urfave/cli/v3"
)

var Edit = &cli.Command{
	Name:   "edit",
	Action: stubAction,
}
