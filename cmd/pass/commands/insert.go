package commands

import (
	"github.com/urfave/cli/v3"
)

var Insert = &cli.Command{
	Name:   "insert",
	Action: stubAction,
}
