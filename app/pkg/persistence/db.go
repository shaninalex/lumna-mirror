package persistence

import (
	"fmt"
	"log"
	"os"
	"time"

	"gitlab.com/shaninalex/lumna/app/models"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func ProvideDB(config *config.Config) *gorm.DB {
	l := MakeLogger(config)
	return connect(config.Database, l)
}

func ApplyMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(
		// users and permissions
		&models.Identity{},
		&models.Credential{},
		&models.RefreshToken{},
		&models.Invitation{},

		// application
		&models.Project{},
		&models.Board{},
		&models.Column{},
		&models.Task{},

		// monitoring and background tasks
		&models.Job{},
		&models.ActivityLog{},
	)

	if err != nil {
		return err
	}

	log.Printf("Database migrated")

	return nil
}

func connect(conf config.Database, l logger.Interface) *gorm.DB {
	var db *gorm.DB

	if config.DetectDatabase(conf) == config.DatabaseTypeSQLite {
		db = connectSqlite(conf, l)
	} else if config.DetectDatabase(conf) == config.DatabaseTypePostgres {
		db = connectPG(conf, l)
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

func connectSqlite(conf config.Database, l logger.Interface) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(conf.SQlite.Url),
		&gorm.Config{Logger: l},
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

func connectPG(conf config.Database, l logger.Interface) *gorm.DB {
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
		&gorm.Config{Logger: l},
	)

	if err != nil {
		panic("failed to connect postgres database: " + err.Error())
	}
	return db
}

func MakeLogger(conf *config.Config) logger.Interface {
	return logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,   // Slow SQL threshold
			LogLevel:                  logger.Silent, // Log level
			IgnoreRecordNotFoundError: true,          // Ignore ErrRecordNotFound error for logger
			ParameterizedQueries:      true,          // Don't include params in the SQL log
			Colorful:                  false,         // Disable color
		},
	)
}
