package database

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib"
)

type DbService struct {
	DbPool *sql.DB
}

var (
	database = os.Getenv("POSTGRES_DB")
	password = os.Getenv("POSTGRES_PASSWORD")
	username = os.Getenv("POSTGRES_USER")
	port     = os.Getenv("DB_PORT")
	host     = os.Getenv("DB_HOST")
	ssl      = os.Getenv("DB_SSL")
)

func Connect() (*DbService, error) {
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", username, password, host, port, database, ssl)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		return nil, err
	}
	return &DbService{DbPool: db}, nil
}

func (s *DbService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	err := s.DbPool.PingContext(ctx)
	if err != nil {
		log.Fatalf("db down: %v", err)
	}

	return map[string]string{
		"message": "It's healthy",
	}
}
