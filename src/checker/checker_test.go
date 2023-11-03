package checker

import (
	"soupdevsolutions/healthchecker/domain"
	"testing"
)

func TestStartStop(t *testing.T) {

	testTarget := domain.NewHealthcheckTarget("Test", "http://localhost:8080")

	healthchecker := Checker{
		SecondsBetweenRuns: 5,
		Targets:            []domain.HealthcheckTarget{testTarget},
	}

	healthchecker.Start()
	// The healthchecker should be running after starting it
	if !healthchecker.IsRunning() {
		t.Error("Expected healthchecker to be running")
	}

	healthchecker.Stop()
	// The healthchecker should not be running after stopping it
	if healthchecker.IsRunning() {
		t.Error("Expected healthchecker to be stopped")
	}

	// Since the checker has a 5 seconds delay, no healthchecks should have been run
	if len(testTarget.Healthchecks) > 0 {
		t.Error("Expected no healthchecks to be run")
	}
}
