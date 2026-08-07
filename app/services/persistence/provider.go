package persistence

import (
	"log"
	"os"

	// _ "github.com/lib/pq"
	// _ "github.com/mattn/go-sqlite3"
	"gitlab.com/shaninalex/lumna/app/pkg/config"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DatabaseType string

var (
	DatabaseTypeSQLite   DatabaseType = "sqlite"
	DatabaseTypePostgres DatabaseType = "postgres"
)

func ProvideDB(conf *config.Config) *gorm.DB {
	return connect(conf)
}

func connect(conf *config.Config) *gorm.DB {
	var db *gorm.DB

	typ := DatabaseType(conf.String("database.type"))
	switch typ {
	case DatabaseTypeSQLite:
		db = connectSqlite(conf)
	case DatabaseTypePostgres:
		db = connectPG(conf)
	default:
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

func connectSqlite(conf *config.Config) *gorm.DB {
	db, err := gorm.Open(
		sqlite.Open(conf.String("database.url")),
		connectionOptions(conf),
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

func connectPG(conf *config.Config) *gorm.DB {
	// dsn := fmt.Sprintf(
	// 	"host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
	// 	conf.Postgres.Host,
	// 	conf.Postgres.Port,
	// 	conf.Postgres.User,
	// 	conf.Postgres.Password,
	// 	conf.Postgres.Database,
	// )

	db, err := gorm.Open(
		postgres.Open(conf.String("database.url")),
		connectionOptions(conf),
	)

	if err != nil {
		panic("failed to connect postgres database: " + err.Error())
	}
	return db
}

func connectionOptions(conf *config.Config) *gorm.Config {
	opt := &gorm.Config{}
	//if conf.Env() != config.EnvironmentDev && conf.Env() != config.EnvironmentTest {
	//	opt.Logger = silentLogger()
	//}
	return opt
}

func silentLogger() logger.Interface {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			LogLevel: logger.Silent,
		},
	)

	return newLogger
}
