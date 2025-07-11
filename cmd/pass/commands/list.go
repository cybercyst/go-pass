package commands

import (
	"github.com/urfave/cli/v3"
)

var DefaultCommand = List.Name

var List = &cli.Command{
	Name:    "ls",
	Aliases: []string{"list"},
	Usage:   "List passwords.",
	Description: `
List names of passwords inside the tree at subfolder by using the tree(1) program. This command is alternatively named list.
`,
	Action: stubAction,
}
