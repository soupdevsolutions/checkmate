package main

import (
	"context"
	"fmt"
	"log"

	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/healthcheck"
	"soupdevsolutions/healthchecker/runner"
)

var db *database.Database
var hcRunner runner.HealthcheckRunner

func initConfig() *config.Config {
	log.Println("reading config")
	config, err := config.ReadConfig()
	if err != nil {
		log.Println("error reading config")
		panic(err)
	}
	return config
}

func initDatabase(ctx context.Context, config *config.Config) *database.Database {
	db, err := database.InitDatabase(ctx, config.Database)
	if err != nil {
		log.Println("error initializing database")
		panic(err)
	}

	var targets []healthcheck.HealthcheckTarget
	for _, target := range config.Runner.Targets {
		targets = append(targets, healthcheck.HealthcheckTarget{
			Name: target.Name,
			Uri:  target.Uri,
		})
	}
	db.Seed(ctx, targets)

	return db
}

func main() {
	log.Println("starting application")
	ctx := context.Background()

	config := initConfig()
	db = initDatabase(ctx, config)

	hcRunner = runner.NewHealthcheckRunner(config.Runner, db, runner.CheckHttpTarget)
	hcRunner.Start()

	log.Println("starting web server")
	router := initRouter()
	router.Run(fmt.Sprintf("%s:%d", config.App.Host, config.App.Port))
}
