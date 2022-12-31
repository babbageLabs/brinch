package bin

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"os"
	"strings"
)

type AppConfig struct {
	Name string
	Env  string
}

type Config struct {
	DB struct {
		URL              string   `yaml:"url"`
		Engine           string   `yaml:"engine"`
		Scripts          string   `yaml:"scripts"`
		FileMatchPattern string   `yaml:"fileMatchPattern"`
		SeedMode         SeedMode `yaml:"seedMode"`
	} `yaml:"db"`
	Logging struct {
		Level logrus.Level `yaml:"level"`
	} `yaml:"logging"`
	App AppConfig `yaml:"app"`
}

func (config *Config) ReadConfig(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		if !config.App.IsTest() {
			return false, err
		}
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil && !config.App.IsTest() {
			Logger.Fatal("error closing file read")
		}
	}(file)

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			if !config.App.IsTest() {
				return false, err
			}
		}
	}

	return true, nil
}

func MustReadConfig(cCtx *cli.Context) Config {
	config := Config{}
	path := cCtx.String(ConfigFlag.Name)
	if path != "" {
		ok, _ := config.ReadConfig(path)
		if ok {
			return config
		}
	}
	Logger.Fatal("Error reading application configuration")

	return Config{}
}

// IsTest return if the application is running in test mode. useful in running tests
func (config *AppConfig) IsTest() bool {
	return strings.ToLower(config.Env) == "test"
}
