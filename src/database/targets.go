package database

import (
	"context"
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
)

func (db *Database) GetTargets(ctx context.Context) ([]healthcheck.HealthcheckTarget, error) {
	rows, err := db.db.QueryContext(ctx, "SELECT id, name, uri FROM targets")
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get targets")
	}
	defer rows.Close()

	targets := make([]healthcheck.HealthcheckTarget, 0)
	for rows.Next() {
		var target healthcheck.HealthcheckTarget
		err := rows.Scan(&target.Id, &target.Name, &target.Uri)
		if err != nil {
			log.Println(err)
			return nil, errors.New("could not get targets")
		}
		targets = append(targets, target)
	}

	return targets, nil
}

func (db *Database) InsertTarget(target *healthcheck.HealthcheckTarget) error {
	_, err := db.db.Exec(
		"INSERT INTO targets (name, uri) VALUES ($1, $2)",
		target.Name,
		target.Uri,
	)
	if err != nil {
		log.Println(err)
		return errors.New("could not insert target")
	}

	return nil
}
