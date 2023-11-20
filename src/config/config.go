package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v2"
)

type DatabaseConfig struct {
	Name     string `yaml:"name"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Host     string `yaml:"host"`
	Port     int    `yaml:"port"`
}

func (db *DatabaseConfig) GetConnectionString() string {
	return fmt.Sprintf("postgresql://%s:%s@%s:%d/%s?sslmode=disable", db.Username, db.Password, db.Host, db.Port, db.Name)
}

type TargetConfig struct {
	Name string `yaml:"name"`
	Uri  string `yaml:"uri"`
}

type RunnerConfig struct {
	Period  int            `yaml:"period"`
	Targets []TargetConfig `yaml:"targets"`
}

type Config struct {
	Database DatabaseConfig `yaml:"database"`
	Runner   RunnerConfig   `yaml:"runner"`
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
