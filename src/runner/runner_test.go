package runner

import (
	"database/sql"
	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/database"
	"soupdevsolutions/healthchecker/healthcheck"
	"testing"
	"time"

	_ "github.com/proullon/ramsql/driver"
)

// a checker that always returns a 200 status code
var ok_checker = func(target healthcheck.HealthcheckTarget) (healthcheck.Healthcheck, error) {
	return healthcheck.Healthcheck{
		Status:    healthcheck.Healthy,
		Timestamp: time.Now().Unix(),
	}, nil
}

func SetupInMemoryDatabase() *database.Database {
	db, err := sql.Open("ramsql", "Test")
	if err != nil {
		panic(err)
	}
	database := database.NewDatabase(db)
	database.Migrate()

	return database
}

func TestStartStop(t *testing.T) {
	db := SetupInMemoryDatabase()

	mockConfig := config.RunnerConfig{
		Period: 5,
	}

	runner := NewHealthcheckRunner(mockConfig, db, ok_checker)

	runner.Start()
	// The healthchecker should be running after starting it
	if !runner.IsRunning() {
		t.Error("expected healthchecker to be running")
	}

	runner.Stop()
	// The healthchecker should not be running after stopping it
	if runner.IsRunning() {
		t.Error("expected healthchecker to be stopped")
	}
}
