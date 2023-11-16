package database

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	log "github.com/sirupsen/logrus"
)

func (d *Database) MigrateDB() error {
	log.Info("Migrate our database")

	driver, err := postgres.WithInstance(d.Client.DB, &postgres.Config{})
	if err != nil {
		return fmt.Errorf("Could not create driver migration: %w", err)
	}

	m, err := migrate.NewWithDatabaseInstance("file:///migrations", "postgres", driver)

	if err != nil {
		log.Error(err.Error())
		return err
	}

	if err := m.Up(); err != nil {
		return fmt.Errorf("Could not up migration: %w", err)

	}

	fmt.Println("Successfully migrate database")
	return nil
}
