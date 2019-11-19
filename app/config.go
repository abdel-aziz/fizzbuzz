package app

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Config struct {
	Port      int
	DebugPort int
	Router    string
}

var config = flag.String("config", "", "Path to the server config file")

func NewConfig() (*Config, error) {
	flag.Parse()

	if config == nil {
		return nil, fmt.Errorf("Config file is required")
	}

	path, err := filepath.Abs(*config)
	if err != nil {
		return nil, err
	}

	file, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var cfg *Config
	err = json.Unmarshal(file, &cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil
}
