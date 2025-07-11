package main

import (
	"context"
	"log"
	"os"

	"github.com/cybercyst/go-pass/cmd/pass/commands"
	"github.com/urfave/cli/v3"
)

func main() {
	cmd := &cli.Command{
		DefaultCommand: commands.DefaultCommand,
		Commands: []*cli.Command{
			commands.Init,
			commands.List,
			commands.Find,
		},
	}

	if err := cmd.Run(context.Background(), os.Args); err != nil {
		log.Fatal(err)
	}
}
