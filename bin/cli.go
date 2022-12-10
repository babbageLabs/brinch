package bin

import "github.com/urfave/cli/v2"

type Commands []*cli.Command
type Flags []cli.Flag

func (comms *Commands) RegisterCmd(cmd *cli.Command) int {
	*comms = append(*comms, cmd)

	return len(*comms)
}

func (flags *Flags) RegisterFlag(flag cli.Flag) int {
	*flags = append(*flags, flag)

	return len(*flags)
}
