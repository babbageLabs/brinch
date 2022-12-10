package cmds

import (
	"github.com/urfave/cli/v2"
)

var SeedCmd = &cli.Command{
	Name:    "seed",
	Aliases: []string{"s"},
	Usage:   "Initialize a database as configured in the config file",
	Action: func(*cli.Context) error {
		//seed := bin.Seed{
		//	path:             "../testdata",
		//	fileMatchPattern: "^expected\\.sql$",
		//	db:               db,
		//}
		return nil
	},
}
