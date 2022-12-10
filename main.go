package main

import (
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/cmds"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"sort"
	"strconv"
)

func main() {
	commands := bin.Commands{}
	flags := bin.Flags{}

	commands.RegisterCmd(cmds.SeedCmd)
	commands.RegisterCmd(cmds.StartCmd)

	// https://cli.urfave.org/v2/examples/flags/
	flags.RegisterFlag(&bin.ConfigFlag)

	app := &cli.App{
		Name:     "brinch",
		Usage:    "Next Gen app development toolkit",
		Flags:    flags,
		Commands: commands,
		Action: func(cCtx *cli.Context) error {
			config := bin.MustReadConfig(cCtx)

			err := os.Setenv("config.Logging.Level", strconv.Itoa(int(config.Logging.Level)))
			if err != nil {
				return err
			}

			return nil
		},
	}

	// https://cli.urfave.org/v2/examples/flags/#ordering
	sort.Sort(cli.FlagsByName(app.Flags))
	sort.Sort(cli.CommandsByName(app.Commands))

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
