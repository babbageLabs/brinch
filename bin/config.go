package bin

import (
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	Db struct {
		Url              string `yaml:"url"`
		Engine           string `yaml:"engine"`
		Scripts          string `yaml:"scripts"`
		FileMatchPattern string `yaml:"fileMatchPattern"`
	} `yaml:"db"`
	Logging struct {
		Level logrus.Level
	} `yaml:"logging"`
}

func (config *Config) ReadConfig(path string) (bool, error) {
	file, err := os.Open(path)
	if err != nil {
		return false, err
	}

	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("error closing file read")
		}
	}(file)

	// Decode YAML file to struct
	if file != nil {
		decoder := yaml.NewDecoder(file)
		if err := decoder.Decode(&config); err != nil {
			return false, err
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
