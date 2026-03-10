package database

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Connect() error {

	databaseURL := os.Getenv("DATABASE_URL")

	if databaseURL == "" {
		return fmt.Errorf("DATABASE_URL is not set")
	}

	databaseName := os.Getenv("DATABASE_NAME")

	if databaseName == "" {
		return fmt.Errorf("DATABASE_NAME is not set")
	}

	config, err := pgxpool.ParseConfig(databaseURL + "/" + databaseName)
	if err != nil {
		return err
	}

	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		return err
	}

	DB = pool

	fmt.Println("Connected to PostgreSQL")

	return nil
}
