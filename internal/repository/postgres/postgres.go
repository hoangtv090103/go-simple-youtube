package postgres

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/lib/pq"

	"gosimpleyoutube/config"
)

type PostgresRepository struct {
	db *sql.DB
}

func NewPostgresRepository(cfg *config.Config) (*sql.DB, error) {
	db, err := sql.Open(
		"postgres",
		fmt.Sprintf(
			"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			cfg.PostgresConfig.Host,
			cfg.PostgresConfig.Port,
			cfg.PostgresConfig.User,
			cfg.PostgresConfig.Password,
			cfg.PostgresConfig.DBName,
		),
	)

	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	// Create database if not exists
	_, err = db.Exec(`CREATE DATABASE youtube`)
	if err != nil {
		if strings.Contains(err.Error(), "database \"youtube\" already exists") {
			return nil, fmt.Errorf("failed to create database: %v", err)
		}
	}

	return db, nil
}

func (r *PostgresRepository) Close() error {
	return r.db.Close()
}
