package bin

import "github.com/urfave/cli/v2"

var ConfigFlag = cli.StringFlag{
	Name:  "config",
	Usage: "Load configuration from `FILE`",
	//Required: true,
	Value:   "./config.yaml",
	Aliases: []string{"c"},
}
