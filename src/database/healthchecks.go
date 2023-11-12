package database

import (
	"context"
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"
)

type HealthchecksRepository struct {
	db *Database
}

func NewHealthchecksRepository(db *Database) HealthchecksRepository {
	return HealthchecksRepository{
		db: db,
	}
}

func (repo *HealthchecksRepository) GetHealthchecks(ctx context.Context, targetId string, limit int) ([]healthcheck.Healthcheck, error) {
	rows, err := repo.db.client.QueryContext(ctx, "SELECT id, status, timestamp FROM healthchecks WHERE target_id = $1 LIMIT $2", targetId, limit)
	if err != nil {
		log.Println(err)
		return nil, errors.New("could not get healthchecks")
	}
	defer rows.Close()

	healthchecks := make([]healthcheck.Healthcheck, 0)
	for rows.Next() {
		var healthcheck healthcheck.Healthcheck
		err := rows.Scan(&healthcheck.Id, &healthcheck.Status, &healthcheck.Timestamp)
		if err != nil {
			log.Println(err)
			return nil, errors.New("could not get healthchecks")
		}
		healthchecks = append(healthchecks, healthcheck)
	}

	return healthchecks, nil
}

func (repo *HealthchecksRepository) InsertHealthcheck(ctx context.Context, target *healthcheck.HealthcheckTarget, healthcheck *healthcheck.Healthcheck) error {
	_, err := repo.db.client.ExecContext(
		ctx,
		"INSERT INTO healthchecks (target_id, status, timestamp) VALUES ($1, $2, $3)",
		target.Id,
		healthcheck.Status,
		healthcheck.Timestamp,
	)
	if err != nil {
		log.Println(err)
		return errors.New("could not insert healthcheck")
	}

	return nil
}
