package runner

import (
	"log"
	"soupdevsolutions/healthchecker/domain"
	"time"
)

type HealthcheckRunner struct {
	Targets []domain.HealthcheckTarget
	Delay   int
	running bool
	checker func(domain.HealthcheckTarget) (domain.Healthcheck, error)
}

func NewHealthcheckRunner(delay int, checker func(domain.HealthcheckTarget) (domain.Healthcheck, error)) HealthcheckRunner {
	return HealthcheckRunner{
		Delay:   delay,
		Targets: []domain.HealthcheckTarget{},
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
			log.Println("Checking target: ", target.Name)

			result, err := c.checker(target)
			if err != nil {
				log.Println("Error checking target: ", err)
				continue
			}

			c.Targets[i].Healthchecks = append(c.Targets[i].Healthchecks, result)
		}
	}
}
