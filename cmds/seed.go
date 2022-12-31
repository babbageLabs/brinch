package cmds

import (
	"github.com/babbageLabs/brinch/bin"
	"github.com/urfave/cli/v2"
)

var SeedCmd = &cli.Command{
	Name:    "seed",
	Aliases: []string{"s"},
	Usage:   "Initialize a database as configured in the config file",
	Action: func(cCtx *cli.Context) error {
		config := bin.MustReadConfig(cCtx)
		seed := bin.Seed{
			Path:             config.DB.Scripts,
			FileMatchPattern: config.DB.FileMatchPattern,
			DB:               bin.MustOpenDBConnection(&config),
			Mode:             config.DB.SeedMode,
		}

		_, err := seed.Seed()
		if err != nil {
			return err
		}
		return nil
	},
}
