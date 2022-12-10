package bin

import (
	"github.com/sirupsen/logrus/hooks/test"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMustOpenDBConnectionSuccess(t *testing.T) {
	config := Config{
		App: AppConfig{
			Name: "brinch",
			Env:  "test",
		},
	}

	logger, hook := test.NewNullLogger()
	Logger = logger

	MustOpenDBConnection(&config)

	assert.Equal(t, 1, len(hook.Entries))
	hook.Reset()
	assert.Nil(t, hook.LastEntry())
}
