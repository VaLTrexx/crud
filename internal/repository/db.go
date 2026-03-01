package repository

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

func newdb() *pgxpool.Pool {
	dsn := os.Getenv("dburl")

	pool, err := pgxpool.New(context.Background(), dsn)
	if err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}
	return pool
}
