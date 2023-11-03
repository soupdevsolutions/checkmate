package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/domain"

	"soupdevsolutions/healthchecker/checker"
)

var healthchecker checker.Checker = checker.Checker{
	SecondsBetweenRuns: 5,
	Targets: []domain.HealthcheckTarget{
		{
			Uri:          "http://www.google.com",
			Name:         "Google",
			Healthchecks: []domain.Healthcheck{},
		},
		{
			Uri:          "http://www.yahoo.com",
			Name:         "Yahoo",
			Healthchecks: []domain.Healthcheck{},
		},
	},
}

func main() {

	router := gin.Default()
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"application": "Healthchecker",
		})
	})

	router.GET("/healthchecks", getHealthchecks)

	healthchecker.Start()
	router.Run("127.0.0.1:8080")
}

func getHealthchecks(c *gin.Context) {

	c.JSON(http.StatusOK, gin.H{"targets": healthchecker.Targets})
}
