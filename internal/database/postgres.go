package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// Dependency injection points
var (
	newPool = pgxpool.New

	sleep = time.Sleep

	fatal = log.Fatal

	ping = func(pool *pgxpool.Pool) error {
		return pool.Ping(context.Background())
	}
)

func Init() {

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres:postgres@127.0.0.1:5433/ai_orchestrator?sslmode=disable"
	}

	var pool *pgxpool.Pool
	var err error

	for i := 1; i <= 10; i++ {

		pool, err = newPool(context.Background(), url)
		if err == nil {

			if err = ping(pool); err == nil {
				DB = pool
				log.Println("✅ PostgreSQL Connected")
				return
			}
		}

		log.Printf("Waiting for PostgreSQL... (%d/10)", i)
		sleep(3 * time.Second)
	}

	fatal("Could not connect to PostgreSQL:", err)
}
