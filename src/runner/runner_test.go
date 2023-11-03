package runner

import (
	"soupdevsolutions/healthchecker/domain"
	"testing"
	"time"
)

func TestStartStop(t *testing.T) {

	runner := HealthcheckRunner{
		Delay:   5,
		Targets: []domain.HealthcheckTarget{domain.NewHealthcheckTarget("Test", "http://localhost:8080")},
	}

	runner.Start()
	// The healthchecker should be running after starting it
	if !runner.IsRunning() {
		t.Error("Expected healthchecker to be running")
	}

	runner.Stop()
	// The healthchecker should not be running after stopping it
	if runner.IsRunning() {
		t.Error("Expected healthchecker to be stopped")
	}

	// Since the checker has a 5 seconds delay, no healthchecks should have been run
	if len(runner.Targets[0].Healthchecks) > 0 {
		t.Error("Expected no healthchecks to be run")
	}
}

func TestRunCheckers(t *testing.T) {
	// a checker that always returns a 200 status code
	ok_checker := func(target domain.HealthcheckTarget) (domain.Healthcheck, error) {
		return domain.Healthcheck{
			StatusCode: 200,
			Timestamp:  time.Now().Unix(),
		}, nil
	}

	runner := NewHealthcheckRunner(1, ok_checker)
	runner.Targets = []domain.HealthcheckTarget{domain.NewHealthcheckTarget("Test", "http://localhost:8080")}

	runner.Start()
	// Give the runner some time to run the healthchecks
	time.Sleep(5 * time.Second)

	if len(runner.Targets[0].Healthchecks) == 0 {
		t.Error("Expected healthchecks to have been run")
	}
}
