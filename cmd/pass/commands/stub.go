package commands

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v3"
)

var stubAction = func(ctx context.Context, c *cli.Command) error {
	fmt.Println(c.Name)
	return nil
}
