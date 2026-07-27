package database

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"

	"checker/internal/config"
)

// Connect создает подключение к PostgreSQL
func Connect(cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s",
		cfg.DBUser,
		cfg.DBPassword,
		cfg.DBHost,
		cfg.DBPort,
		cfg.DBName,
	)

	db, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		db.Close()
		return nil, err
	}

	return db, nil
}

// Close закрывает подключение к PostgreSQL
func Close(db *pgxpool.Pool) {
	if db != nil {
		db.Close()
	}
}
