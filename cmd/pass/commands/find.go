package commands

import (
	"github.com/urfave/cli/v3"
)

var Find = &cli.Command{
	Name:    "find",
	Aliases: []string{"search"},
	Usage:   "List passwords that match pass-names.",
	Description: `
List names of passwords inside the tree that match pass-names by using the tree(1) program. This command is alternatively named search.
`,
	Action: stubAction,
}
