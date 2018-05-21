package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/shootnix/golatin-logger-service/config"
	"github.com/shootnix/golatin-logger-service/daemon"
	"github.com/shootnix/golatin-logger-service/models"
	"log"
)

type Config struct {
	Database DatabaseConfig
}

type DatabaseConfig struct {
	Address string
}

func main() {

	cfg := config.Load("config.toml")

	db, err := OpenDatabaseConnection(cfg.Database)
	if err != nil {
		log.Fatal(err.Error())
	}
	models.Pg = db

	daemon := daemon.NewDaemon(cfg.Daemon)

	log.Println("Starting...")
	daemon.Run()
}

func OpenDatabaseConnection(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	db, err := sqlx.Open("postgres", cfg.ConnectionInfo)
	if err != nil {
		return db, err
	}
	if err = db.Ping(); err != nil {
		return db, err
	}

	return db, err
}
