package cmds

import "github.com/urfave/cli/v2"

var ConfigFlag = cli.StringFlag{
	Name:     "config",
	Usage:    "The path to the config path",
	Required: true,
}
