package runner

import (
	"context"
	"log"
	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/healthcheck"
	"time"
)

type HealthcheckRunner struct {
	Period  int
	running bool
	checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)
	db      *database.Database
	targets []healthcheck.HealthcheckTarget
}

func NewHealthcheckRunner(config config.RunnerConfig, db *database.Database, checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)) HealthcheckRunner {
	return HealthcheckRunner{
		Period:  config.Period,
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
	healthchecksRepo := database.NewHealthchecksRepository(c.db)

	for c.running {
		time.Sleep(time.Duration(c.Period) * time.Second)

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
			err = healthchecksRepo.InsertHealthcheck(ctx, &targets[i], &result)
			if err != nil {
				log.Println("Error inserting healthcheck: ", err)
				continue
			}
		}
		c.targets = targets
	}
}

func (c *HealthcheckRunner) Targets() []healthcheck.HealthcheckTarget {
	return c.targets
}
