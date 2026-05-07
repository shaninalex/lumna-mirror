package persistence

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/database/sqlite3"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/sqlite3"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/golang-migrate/migrate/v4/source/iofs"

	_ "github.com/mattn/go-sqlite3"

	"gitlab.com/shaninalex/lumna"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
)

func ApplyMigrations(c *config.Config) error {
	switch config.DetectDatabase(c.Database) {
	case config.DatabaseTypePostgres:
		return MigratePostgres(ProvideDB(c))
	case config.DatabaseTypeSQLite:
		return MigrateSQLite(ProvideDB(c))
	default:
		panic("driver not supported")
	}
}

func MigratePostgres(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	dr, err := postgres.WithInstance(sqlDB, &postgres.Config{})
	if err != nil {
		return err
	}
	defer dr.Close()

	d, err := iofs.New(lumna.StaticFS("resources/migrations/postgres"), ".")
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		d,
		"postgres",
		dr,
	)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("Postgres database migrated")
	return nil
}

func MigrateSQLite(db *gorm.DB) error {
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}
	sourceDriver, err := iofs.New(lumna.StaticFS("resources/migrations/sqlite"), ".")
	if err != nil {
		return err
	}

	dbDriver, err := sqlite3.WithInstance(sqlDB, &sqlite3.Config{})
	if err != nil {
		return err
	}

	m, err := migrate.NewWithInstance(
		"iofs",
		sourceDriver,
		"sqlite3",
		dbDriver,
	)
	if err != nil {
		return err
	}

	err = m.Up()

	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	fmt.Println("SQLite database migrated")
	return nil
}
