package checker

import (
	"log"
	"soupdevsolutions/healthchecker/domain"
	"time"
)

type Checker struct {
	Targets            []domain.HealthcheckTarget
	SecondsBetweenRuns int
	running            bool
}

func (c *Checker) Start() {
	c.running = true
	go c.Run()
}

func (c *Checker) Stop() {
	c.running = false
}

func (c *Checker) IsRunning() bool {
	return c.running
}

func (c *Checker) Run() {
	for c.running {
		time.Sleep(time.Duration(c.SecondsBetweenRuns) * time.Second)

		for i, target := range c.Targets {
			log.Println("Checking target: ", target.Name)

			result := domain.Healthcheck{
				StatusCode: 200,
				Timestamp:  time.Now().Unix(),
			}

			c.Targets[i].Healthchecks = append(c.Targets[i].Healthchecks, result)
		}

	}
}
