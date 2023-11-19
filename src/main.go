package main

import (
	"context"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"

	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/runner"
	"soupdevsolutions/healthchecker/viewmodels"
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
		views.GET("/", getIndexView)
		views.GET("/targets", getTargetsView)
		views.GET("/target", getTargetView)
	}

	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"application": "Healthchecker",
			})
		})
		api.GET("/healthchecks", getTargetsJson)
	}

	router.Run("127.0.0.1:8080")
}

func getIndexView(c *gin.Context) {
	tmpl := template.Must(template.ParseFiles("../templates/index.html"))
	err := tmpl.Execute(c.Writer, nil)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func getTargetsJson(c *gin.Context) {
	repo := database.NewTargetsRepository(db)
	targets, err := repo.GetTargets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"targets": targets})
}

func getTargetsView(c *gin.Context) {
	repo := database.NewTargetsRepository(db)
	targets, err := repo.GetTargets(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var targetsVms []viewmodels.TargetViewModel

	for _, target := range targets {
		targetsVms = append(targetsVms, viewmodels.NewTargetViewModel(&target))
	}

	tmpl := template.Must(template.ParseFiles("../templates/targets_list.html"))

	fmt.Printf("targets: %+v", targetsVms)

	err = tmpl.Execute(c.Writer, gin.H{"targets": targetsVms})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}

func getTargetView(c *gin.Context) {
	targetId := c.Query("id")
	repo := database.NewTargetsRepository(db)

	target, err := repo.GetTarget(c, targetId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	targetVm := viewmodels.NewTargetViewModel(target)

	tmpl := template.Must(template.ParseFiles("../templates/target.html"))
	err = tmpl.Execute(c.Writer, gin.H{"target": targetVm})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
}
