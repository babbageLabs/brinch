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
			Path:             config.Db.Scripts,
			FileMatchPattern: config.Db.FileMatchPattern,
			Db:               bin.MustOpenDbConnection(&config),
		}

		_, err := seed.Seed()
		if err != nil {
			return err
		}
		return nil
	},
}
