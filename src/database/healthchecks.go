package database

import (
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
)

func (db *Database) InsertHealthcheck(target *healthcheck.HealthcheckTarget, healthcheck *healthcheck.Healthcheck) error {
	_, err := db.db.Exec(
		"INSERT INTO healthchecks (target_id, status, timestamp) VALUES ($1, $2, $3)",
		target.Id,
		healthcheck.Status,
		healthcheck.Timestamp,
	)
	if err != nil {
		log.Fatal(err)
		return errors.New("could not insert healthcheck")
	}

	return nil
}
