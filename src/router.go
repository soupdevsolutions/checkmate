package main

import (
	"fmt"
	"html/template"
	"net/http"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/viewmodels"

	"github.com/gin-gonic/gin"
)

func initRouter() *gin.Engine {
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

	return router
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
