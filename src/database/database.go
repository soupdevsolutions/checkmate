package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
)

type Database struct {
	db *sql.DB
}

func Connect(ctx context.Context) (*Database, error) {
	connectionInfo := ""
	db, err := sql.Open("postgres", connectionInfo)
	if err != nil {
		log.Fatal(err)
		return nil, errors.New("could not open a connection to the database")
	}

	if err := db.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil, errors.New("could not ping the database")
	}

	return &Database{db: db}, nil
}

func (db *Database) InsertTarget(target *healthcheck.HealthcheckTarget) error {
	_, err := db.db.Exec(
		"INSERT INTO targets (target_id, name, uri) VALUES ($1, $2, $3)",
		target.Id,
		target.Name,
		target.Uri,
	)
	if err != nil {
		log.Fatal(err)
		return errors.New("could not insert target")
	}

	return nil
}

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
