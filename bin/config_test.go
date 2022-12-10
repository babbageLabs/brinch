package bin

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestConfig_ReadConfig(t *testing.T) {
	config := Config{
		App: AppConfig{
			Name: "",
			Env:  "test",
		},
	}
	readConfig, err := config.ReadConfig("../testData/config.yaml")
	assert.NoError(t, err)
	assert.Equal(t, true, readConfig)

}

func TestAppConfig_IsTest(t *testing.T) {
	appConfig := AppConfig{
		Name: "brinch",
		Env:  "test",
	}

	config := Config{
		App: appConfig,
	}
	assert.Equal(t, config.App.IsTest(), true)
}
