package database

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"
	"soupdevsolutions/healthchecker/config"
	"soupdevsolutions/healthchecker/healthcheck"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

type Database struct {
	db     *sql.DB
	config config.DatabaseConfig
}

func NewDatabase(db *sql.DB) *Database {
	return &Database{db: db}
}

func InitDatabase(ctx context.Context, config config.DatabaseConfig) (*Database, error) {
	log.Println("initializing database")
	database, err := Connect(ctx, config.GetConnectionString())
	if err != nil {
		log.Fatal("error connecting to database")
	}
	err = database.Migrate()
	if err != nil {
		log.Fatal("error applying migrations")
	}

	return database, nil
}

func Connect(ctx context.Context, connectionString string) (*Database, error) {
	log.Println("connecting to database")
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		log.Printf("error opening connection to database: %v", err)
		return nil, errors.New("could not open a connection to the database")
	}

	if err := db.PingContext(ctx); err != nil {
		log.Printf("error pinging database: %v", err)
		return nil, errors.New("could not ping the database")
	}

	return &Database{db: db}, nil
}

func (db *Database) Migrate() error {
	log.Println("applying migrations")
	driver, err := postgres.WithInstance(db.db, &postgres.Config{MigrationsTable: "migrations"})
	if err != nil {
		log.Println(err)
		return errors.New("could not create a postgres instance")
	}
	fmt.Println(db.config.Name)
	m, err := migrate.NewWithDatabaseInstance(
		"file://../migrations",
		db.config.Name, driver)
	if err != nil {
		log.Println(err)
		return errors.New("could not init `migrate`")
	}

	if err = m.Up(); err != nil {
		log.Println(err)
		return errors.New("could not apply migrations")
	}

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

	targetsRepo := NewTargetsRepository(db)
	for _, target := range targets {
		targetsRepo.InsertTarget(&target)
	}
}
