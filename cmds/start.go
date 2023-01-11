package cmds

import (
	"github.com/babbageLabs/brinch/bin"
	"github.com/babbageLabs/brinch/bin/app"
	"github.com/urfave/cli/v2"
)

var StartCmd = &cli.Command{
	Name:    "start",
	Aliases: []string{},
	Usage:   "start an instance",
	Action: func(cCtx *cli.Context) error {
		config := bin.MustReadConfig(cCtx)

		instance := app.Instance{
			Name:    config.App.Name,
			Address: config.App.Address,
			Routes:  nil,
		}

		err := instance.Start()
		if err != nil {
			return err
		}

		return nil
	},
}
