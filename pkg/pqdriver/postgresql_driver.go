package pqdriver

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"time"

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
	MaxLifetime  string
	MaxIdleConns string
	MaxOpenConns string
}

type postgresDB struct {
	Conf Config
}

func (db *postgresDB) Connect() *sqlx.DB {
	if db.Conf.SSLMode == "" {
		db.Conf.SSLMode = SSLModeDisable
	}
	dsName := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		db.Conf.Host, db.Conf.Port, db.Conf.User, db.Conf.Pass, db.Conf.DatabaseName, db.Conf.SSLMode)
	conn, err := sqlx.Connect("postgres", dsName)
	maxOpenConns, _ := strconv.Atoi(db.Conf.MaxOpenConns)
	maxIdleConns, _ := strconv.Atoi(db.Conf.MaxIdleConns)
	maxLifetime, _ := strconv.Atoi(db.Conf.MaxLifetime)
	if maxOpenConns > 0 {
		conn.SetMaxOpenConns(maxOpenConns) // The default is 0 (unlimited), ex: 1000
	}
	if maxIdleConns > 0 {
		conn.SetMaxIdleConns(maxIdleConns) // The default maxIdleConns = 2, ex: 10
	}
	conn.SetConnMaxLifetime(time.Duration(maxLifetime)) // MaxLifetime = 0, Connections are reused forever
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
		// The default SSL mode is "disable", ex: "verify-full"
		SSLMode: os.Getenv("POSTGRES_SSL_MODE"),
		// The default maxLifetime = 0, Connections are reused forever, ex: "60"
		MaxLifetime: os.Getenv("POSTGRES_MAX_LIFETIME"),
		// The default maxIdleConns = 2, ex: 10
		MaxIdleConns: os.Getenv("POSTGRES_MAX_IDLE_CONNS"),
		// The default is 0 (unlimited), ex: 1000
		MaxOpenConns: os.Getenv("POSTGRES_MAX_OPEN_CONNS"),
	}
}
