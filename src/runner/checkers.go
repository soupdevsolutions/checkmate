package runner

import (
	"fmt"
	"net/http"
	"soupdevsolutions/healthchecker/healthcheck"
	"time"
)

func CheckHttpTarget(target healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error) {
	fmt.Println("Checking HTTP target: ", target.Name)

	resp, err := http.Get(target.Uri)
	if err != nil {
		return healthcheck.Healthcheck{}, err
	}
	return healthcheck.Healthcheck{
		Status:    healthcheck.FromStatusCode(resp.StatusCode),
		Timestamp: time.Now().Unix(),
	}, nil
}
