package main

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/runner"
)

var db *database.Database
var healthchecker runner.HealthcheckRunner

func main() {
	log.Println("starting application")

	ctx := context.Background()

	log.Println("reading config")
	config, err := config.ReadConfig()
	if err != nil {
		log.Println("error reading config")
		panic(err)
	}

	db, err = database.InitDatabase(ctx, config.Database)
	if err != nil {
		log.Println("error initializing database")
		panic(err)
	}
	db.Seed()

	healthchecker = runner.NewHealthcheckRunner(5, db, runner.CheckHttpTarget)
	healthchecker.Start()

	log.Println("starting web server")
	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "Healthchecker",
		})
	})

	router.GET("/healthchecks", getHealthchecks)
	router.Run("127.0.0.1:8080")
}

func getHealthchecks(c *gin.Context) {
	targets := healthchecker.Targets()
	c.JSON(http.StatusOK, gin.H{"targets": targets})
}
