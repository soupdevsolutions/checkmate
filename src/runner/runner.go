package runner

import (
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
	"time"
)

type HealthcheckRunner struct {
	Targets []healthcheck.HealthcheckTarget
	Delay   int
	running bool
	checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)
}

func NewHealthcheckRunner(delay int, checker func(healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error)) HealthcheckRunner {
	return HealthcheckRunner{
		Delay:   delay,
		Targets: []healthcheck.HealthcheckTarget{},
		running: false,
		checker: checker,
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
	for c.running {
		time.Sleep(time.Duration(c.Delay) * time.Second)

		for i, target := range c.Targets {
			result, err := c.checker(target)
			if err != nil {
				log.Println("Error checking target: ", err)
				continue
			}

			c.Targets[i].Healthchecks = append(c.Targets[i].Healthchecks, result)
		}
	}
}
