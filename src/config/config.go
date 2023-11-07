package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Name string `yaml:"name"`
	Host string `yaml:"host"`
	Port int    `yaml:"port"`
}

func (db *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("Host=%s; Port=%d; Database=%s;", db.Host, db.Port, db.Name)
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
}

func ReadConfig() (*Config, error) {
	environment := GetEnvironment()

	f, err := os.Open(fmt.Sprintf("../config/%s.yaml", environment))
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var config Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
