package bin

import (
	"github.com/stretchr/testify/assert"
	"github.com/urfave/cli/v2"
	"testing"
)

func TestCommands_RegisterCmd(t *testing.T) {
	cmd := &cli.Command{}
	cmds := Commands{}

	count := cmds.RegisterCmd(cmd)

	assert.Equal(t, len(cmds), count)
}

func TestFlags_RegisterFlag(t *testing.T) {
	flags := Flags{}

	count := flags.RegisterFlag(&ConfigFlag)
	assert.Equal(t, len(flags), count)
}
