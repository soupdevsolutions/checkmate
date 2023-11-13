package main

import (
	"context"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/runner"
)

var db *database.Database
var hcRunner runner.HealthcheckRunner

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
	db.Seed(ctx)

	hcRunner = runner.NewHealthcheckRunner(5, db, runner.CheckHttpTarget)
	hcRunner.Start()

	log.Println("starting web server")
	router := gin.Default()

	views := router.Group("/")
	{
		views.GET("/healthchecks", getHealthchecksView)
	}

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"application": "Healthchecker",
			})
		})
		api.GET("/healthchecks", getHealthchecksJson)
	}

	router.Run("127.0.0.1:8080")
}

func getHealthchecksJson(c *gin.Context) {
	repo := database.NewTargetsRepository(db)
	targets, err := repo.GetTargets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"targets": targets})
}

func getHealthchecksView(c *gin.Context) {
	repo := database.NewTargetsRepository(db)
	targets, err := repo.GetTargets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	tmpl := template.Must(template.ParseFiles("../templates/index.html"))

	err = tmpl.Execute(c.Writer, gin.H{"targets": targets})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
