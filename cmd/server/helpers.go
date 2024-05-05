package main

import (
	"errors"
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	AWSAccessKeyID     string `yaml:"aws_access_key_id"`
	AWSSecretAccessKey string `yaml:"aws_secret_access_key"`
	AWSRegion          string `yaml:"aws_region"`
}

func ParseConfig() (*Config, error) {
	configFile := os.Getenv("CONFIG_FILE")
	if configFile == "" {
		return nil, errors.New("CONFIG_FILE not specified")
	}
	configSrc, err := os.ReadFile(configFile)
	if err != nil {
		return nil, fmt.Errorf("could not decode file at %s: %w", configFile, err)
	}

	cfg := new(Config)
	if err = yaml.Unmarshal(configSrc, &cfg); err != nil {
		return nil, fmt.Errorf("could not parse yaml at %s: %w", configFile, err)
	}
	return cfg, nil
}
