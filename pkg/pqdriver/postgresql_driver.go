package pqdriver

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

const (
	DefaultPort       = "5432"
	SSLModeVerifyFull = "verify-full"
	SSLModeDisable    = "disable"
	SSLModeRequire    = "require"
)

// PostgreSQLDriver is the interface
type PostgreSQLDriver interface {
	Connect() *sqlx.DB
}

// Config is a model for connect PosgreSQL
type Config struct {
	User         string
	Pass         string
	Host         string
	DatabaseName string
	Port         string
	SSLMode      string
}

type postgresDB struct {
	Conf Config
}

func (db *postgresDB) Connect() *sqlx.DB {
	dsName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Conf.Host, db.Conf.Port, db.Conf.User, db.Conf.Pass, db.Conf.DatabaseName, db.Conf.SSLMode)
	conn, err := sqlx.Connect("postgres", dsName)
	if err != nil {
		log.Fatalln(err)
	} else {
		log.Println("PostgreSQL Connected!")
	}
	return conn
}

// New for create PostgresSQL driver
func New(config Config) PostgreSQLDriver {
	return &postgresDB{
		Conf: config,
	}
}

// ConfigEnv for create Config by Env
func ConfigEnv() Config {
	return Config{
		User:         os.Getenv("POSTGRES_USER"),
		Pass:         os.Getenv("POSTGRES_PASS"),
		Host:         os.Getenv("POSTGRES_HOST"),
		DatabaseName: os.Getenv("POSTGRES_DATABASE"),
		Port:         os.Getenv("POSTGRES_PORT"),
		SSLMode:      os.Getenv("POSTGRES_SSL_MODE"),
	}
}
