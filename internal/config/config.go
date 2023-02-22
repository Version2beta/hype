package internal

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Site struct {
		Title string `yaml:"title"`
		URL   string `yaml:"url"`
	} `yaml:"site"`
	Author []struct {
		Name  string `yaml:"name"`
		Email string `yaml:"email"`
		URL   string `yaml:"url"`
		Bio   string `yaml:"bio"`
	} `yaml:"author"`
}

func LoadConfig(path string) ([]byte, error) {
	// Read YAML file
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func GetConfig(cfgData []byte) (*Config, error) {
	// Parse YAML data into Config struct
	var config Config
	if err := yaml.Unmarshal(cfgData, &config); err != nil {
		return nil, err
	}

	return &config, nil
}
