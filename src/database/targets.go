package database

import (
	"context"
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
)

type TargetsRepository struct {
	db *Database
}

func NewTargetsRepository(db *Database) TargetsRepository {
	return TargetsRepository{
		db: db,
	}
}

func (repo *TargetsRepository) GetTarget(ctx context.Context, id string) (*healthcheck.HealthcheckTarget, error) {

	tx, err := repo.db.client.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get targets")
	}

	row := tx.QueryRowContext(ctx, "SELECT id, name, uri FROM targets WHERE id = $1", id)

	var target healthcheck.HealthcheckTarget
	err = row.Scan(&target.Id, &target.Name, &target.Uri)
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get target")
	}

	healthchecksLimit := 5
	healthchecksRows, err := tx.QueryContext(
		ctx,
		"SELECT id, status, timestamp FROM healthchecks WHERE target_id = $1 ORDER BY timestamp DESC LIMIT $2",
		target.Id,
		healthchecksLimit)
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get healthchecks for target " + target.Name)
	}
	defer healthchecksRows.Close()

	for healthchecksRows.Next() {
		var healthcheck healthcheck.Healthcheck
		err := healthchecksRows.Scan(&healthcheck.Id, &healthcheck.Status, &healthcheck.Timestamp)
		if err != nil {
			log.Println(err)
			return nil, errors.New("could not get healthchecks for target " + target.Name)
		}
		target.Healthchecks = append(target.Healthchecks, healthcheck)
	}

	return &target, nil
}

func (repo *TargetsRepository) GetTargets(ctx context.Context) ([]healthcheck.HealthcheckTarget, error) {

	tx, err := repo.db.client.BeginTx(ctx, nil)
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get targets")
	}
	defer tx.Rollback()

	targetsRows, err := tx.QueryContext(ctx, "SELECT id, name, uri FROM targets")
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get targets")
	}
	defer targetsRows.Close()

	targets := make([]healthcheck.HealthcheckTarget, 0)
	for targetsRows.Next() {
		var target healthcheck.HealthcheckTarget
		err := targetsRows.Scan(&target.Id, &target.Name, &target.Uri)
		if err != nil {
			log.Println(err)
			return nil, errors.New("could not get targets")
		}
		targets = append(targets, target)
	}

	for i, target := range targets {
		healthchecksLimit := 5
		healthchecksRows, err := tx.QueryContext(
			ctx,
			"SELECT id, status, timestamp FROM healthchecks WHERE target_id = $1 ORDER BY timestamp DESC LIMIT $2",
			target.Id,
			healthchecksLimit)
		if err != nil {
			log.Println(err)
			return nil, errors.New("could not get healthchecks for target " + target.Name)
		}
		defer healthchecksRows.Close()

		for healthchecksRows.Next() {
			var healthcheck healthcheck.Healthcheck
			err := healthchecksRows.Scan(&healthcheck.Id, &healthcheck.Status, &healthcheck.Timestamp)
			if err != nil {
				log.Println(err)
				return nil, errors.New("could not get healthchecks for target " + target.Name)
			}
			targets[i].Healthchecks = append(targets[i].Healthchecks, healthcheck)
		}
	}

	return targets, nil
}

func (repo *TargetsRepository) InsertTarget(ctx context.Context, target *healthcheck.HealthcheckTarget) error {
	_, err := repo.db.client.ExecContext(
		ctx,
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
