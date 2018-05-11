package main

import (
	"io/ioutil"

	"gopkg.in/yaml.v2"

	"github.com/iwaaya/roomo/db"
	"github.com/iwaaya/roomo/obs"
)

type Config struct {
	DB  db.Config  `yaml:db`
	OBS obs.Config `yaml:obs`
}

func NewConfig(file string) (*Config, error) {
	buf, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	var c Config
	err = yaml.Unmarshal(buf, &c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}
