package commands

import (
	"context"
	"log"

	"github.com/urfave/cli/v3"
)

var stubAction = func(ctx context.Context, c *cli.Command) error {
	log.Fatalf("%s is not implemented!", c.Name)
	return nil
}
