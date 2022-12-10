package bin

import (
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Config struct {
	db struct {
		Url              string `yaml:"url"`
		Engine           string `yaml:"engine"`
		Scripts          string `yaml:"scripts"`
		FileMatchPattern string `yaml:"fileMatchPattern"`
	} `yaml:"db"`
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
