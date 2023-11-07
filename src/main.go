package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/healthcheck"
	"soupdevsolutions/healthchecker/runner"
)

var healthchecker runner.HealthcheckRunner = runner.NewHealthcheckRunner(5, runner.CheckHttpTarget)

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "Healthchecker",
		})
	})

	router.GET("/healthchecks", getHealthchecks)

	healthchecker.Targets = []healthcheck.HealthcheckTarget{
		{
			Uri:          "http://www.google.com",
			Name:         "Google",
			Healthchecks: []healthcheck.Healthcheck{},
		},
		{
			Uri:          "http://www.yahoo.com",
			Name:         "Yahoo",
			Healthchecks: []healthcheck.Healthcheck{},
		},
	}
	healthchecker.Start()
	router.Run("127.0.0.1:8080")
}

func getHealthchecks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"targets": healthchecker.Targets})
}
