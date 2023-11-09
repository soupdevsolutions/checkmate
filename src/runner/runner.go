package runner

import (
	"context"
	"log"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/healthcheck"
	"time"
)

type HealthcheckRunner struct {
	Delay   int
	running bool
	checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)
	db      *database.Database
	targets []healthcheck.HealthcheckTarget
}

func NewHealthcheckRunner(delay int, db *database.Database, checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)) HealthcheckRunner {
	return HealthcheckRunner{
		Delay:   delay,
		running: false,
		checker: checker,
		db:      db,
	}
}

func (c *HealthcheckRunner) Start() {
	c.running = true
	go c.run()
}

func (c *HealthcheckRunner) Stop() {
	c.running = false
}

func (c *HealthcheckRunner) IsRunning() bool {
	return c.running
}

func (c *HealthcheckRunner) run() {
	ctx := context.Background()
	targetsRepo := database.NewTargetsRepository(c.db)

	for c.running {
		time.Sleep(time.Duration(c.Delay) * time.Second)

		targets, err := targetsRepo.GetTargets(ctx)
		if err != nil {
			log.Println("error getting targets: ", err)
			continue
		}

		for i, target := range targets {
			result, err := c.checker(target)
			if err != nil {
				log.Println("Error checking target: ", err)
				continue
			}

			targets[i].Healthchecks = append(targets[i].Healthchecks, result)
		}
		c.targets = targets
	}
}

func (c *HealthcheckRunner) Targets() []healthcheck.HealthcheckTarget {
	return c.targets
}
