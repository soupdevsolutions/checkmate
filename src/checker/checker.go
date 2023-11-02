package checker

import (
	"fmt"
	"soupdevsolutions/healthchecker/domain"
	"time"
)

type Checker struct {
	Targets           []domain.HealthcheckTarget
	SecondBetweenRuns int
}

func (c *Checker) Check() {
	for {
		for i, target := range c.Targets {
			fmt.Println("Checking target: ", target.Name)

			result := domain.Healthcheck{
				StatusCode: 200,
				Timestamp:  time.Now().Unix(),
			}

			c.Targets[i].Healthchecks = append(c.Targets[i].Healthchecks, result)

			fmt.Println("Target: ", len(c.Targets[i].Healthchecks))
		}

		time.Sleep(time.Duration(c.SecondBetweenRuns) * time.Second)
	}
}
