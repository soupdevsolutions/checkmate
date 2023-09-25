package main

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/domain"
)

func main() {
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
	healthchecks := []domain.Healthcheck{
		{
			Uri:    "google.com",
			Name:   "Google",
			Status: "OK",
		},
		{
			Uri:    "facebook.com",
			Name:   "Facebook",
			Status: "OK",
		},
		{
			Uri:    "twitter.com",
			Name:   "Twitter",
			Status: "OK",
		},
	}

	c.JSON(http.StatusOK, healthchecks)
}
