package postgres

import (
	"todo/config"
	"database/sql"
	"fmt"
	loggerPackage "todo/pkg/logger"
)

func createConnectionString (cfg config.PostgresConfig) string {
	return fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s search_path=%s aalmode=disable",
		cfg.Host, cfg.Port, cfg.User,
		cfg.Password, cfg.DbName, cfg.DbSchema,
	)
}

func NewPostgresClient(cfg config.PostgresConfig) (*sql.DB, error) {
	logger := loggerPackage.Get()
	var err error = nil 
	db, err := sql.Open("postgres", createConnectionString(cfg))
	if err == nil {
		if err = db.Ping(); err == nil {
			return db, nil
		}else {
			logger.Errorf("Failed to connct to db, no ping")
		}
	}else {
		logger.Errorf ("Failed to connect to db")
	}
		return nil, err
}