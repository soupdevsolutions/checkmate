package database

import (
	"context"
	"database/sql"
	"errors"
	"log"
	"soupdevsolutions/healthchecker/healthcheck"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	db *sql.DB
}

func InitDatabase(ctx context.Context, connectionString string) (*Database, error) {
	database, err := Connect(ctx, connectionString)
	if err != nil {
		log.Println("error connecting to database")
		panic(err)
	}
	err = database.Migrate()
	if err != nil {
		log.Println("error applying migrations")
		panic(err)
	}

	return database, nil
}

func Connect(ctx context.Context, connectionString string) (*Database, error) {
	log.Println("connecting to database")
	db, err := sql.Open("postgres", connectionString)
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

func (db *Database) Migrate() error {
	log.Println("applying migrations")
	driver, err := postgres.WithInstance(db.db, &postgres.Config{})
	if err != nil {
		log.Fatal(err)
		return errors.New("could not create a postgres instance")
	}
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		"postgres", driver)
	if err != nil {
		log.Fatal(err)
		return errors.New("could not init `migrate`")
	}

	m.Up()

	return nil
}

func (db *Database) Seed() {
	log.Println("seeding database")
	targets := []healthcheck.HealthcheckTarget{
		{
			Uri:          "http://www.google.com",
			Name:         "Google",
			Healthchecks: []healthcheck.Healthcheck{},
		},
		{
			Uri:          "http://www.yahoo.com",
			Name:         "Yahoo",
			Healthchecks: []healthcheck.Healthcheck{},
		},
	}

	for _, target := range targets {
		db.InsertTarget(&target)
	}
}
