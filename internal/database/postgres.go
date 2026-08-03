package database

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func Init() {

	url := os.Getenv("DATABASE_URL")
	if url == "" {
		url = "postgres://postgres:postgres@127.0.0.1:5433/ai_orchestrator?sslmode=disable"
	}

	var pool *pgxpool.Pool
	var err error

	for i := 1; i <= 10; i++ {

		pool, err = pgxpool.New(context.Background(), url)
		if err == nil {

			if err = pool.Ping(context.Background()); err == nil {
				DB = pool
				log.Println("✅ PostgreSQL Connected")
				return
			}
		}

		log.Printf("Waiting for PostgreSQL... (%d/10)", i)
		time.Sleep(3 * time.Second)
	}

	log.Fatal("Could not connect to PostgreSQL:", err)
}
