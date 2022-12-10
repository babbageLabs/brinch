package main

import (
	"github.com/babbageLabs/brinch/cmds"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {
	commands := cmds.Commands{}
	flags := cmds.Flags{}

	commands.RegisterCmd(cmds.SeedCmd)
	commands.RegisterCmd(cmds.StartCmd)

	flags.RegisterFlag(&cmds.ConfigFlag)

	app := &cli.App{
		Name:     "brinch",
		Usage:    "Next Gen app development toolkit",
		Flags:    flags,
		Commands: commands,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
