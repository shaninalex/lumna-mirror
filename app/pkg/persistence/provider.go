package persistence

import (
	"fmt"

	_ "github.com/lib/pq"
	_ "github.com/mattn/go-sqlite3"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ProvideDB(config *config.Config) *gorm.DB {
	return connect(config.Database)
}

func connect(conf config.Database) *gorm.DB {
	var db *gorm.DB

	if config.DetectDatabase(conf) == config.DatabaseTypeSQLite {
		db = connectSqlite(conf)
	} else if config.DetectDatabase(conf) == config.DatabaseTypePostgres {
		db = connectPG(conf)
	} else {
		panic("database type not supported")
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if err = sqlDB.Ping(); err != nil {
		panic("failed to ping database: " + err.Error())
	}

	return db
}

func connectSqlite(conf config.Database) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(conf.SQlite.Url),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect sqlite database: " + err.Error())
	}

	// Enable foreign keys for SQLite
	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	if _, err := sqlDB.Exec("PRAGMA foreign_keys = ON"); err != nil {
		panic(err)
	}
	return db
}

func connectPG(conf config.Database) *gorm.DB {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		conf.Postgres.Host,
		conf.Postgres.Port,
		conf.Postgres.User,
		conf.Postgres.Password,
		conf.Postgres.Database,
	)

	db, err := gorm.Open(
		postgres.Open(dsn),
		&gorm.Config{},
	)

	if err != nil {
		panic("failed to connect postgres database: " + err.Error())
	}
	return db
}
