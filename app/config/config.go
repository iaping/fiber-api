package config

import (
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

type Config struct {
	Debug bool  `yaml:"Debug"`
	Http  Http  `yaml:"Http"`
	Db    Db    `yaml:"Db"`
	Redis Redis `yaml:"Redis"`
}

func Load(path string) (*Config, error) {
	path, err := filepath.Abs(path)
	if err != nil {
		return nil, err
	}

	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var conf Config
	err = yaml.Unmarshal(data, &conf)
	return &conf, err
}
