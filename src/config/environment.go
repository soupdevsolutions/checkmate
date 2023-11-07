package config

import (
	"os"
	"strings"
)

type Environment int

const (
	Local  Environment = 1
	Remote Environment = 2
)

func (e Environment) String() string {
	switch e {
	case Local:
		return "local"
	case Remote:
		return "remote"
	default:
		return "local"
	}
}

func GetEnvironment() Environment {
	environment := os.Getenv("ENVIRONMENT")

	switch strings.ToLower(environment) {
	case "local":
		return Local
	case "remote":
		return Remote
	default:
		return Local
	}
}
