package database

import (
	"context"
	"errors"
	"os"
	"testing"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func TestInit_Success(t *testing.T) {

	oldNewPool := newPool
	oldPing := ping
	oldSleep := sleep
	oldFatal := fatal
	oldDB := DB

	defer func() {
		newPool = oldNewPool
		ping = oldPing
		sleep = oldSleep
		fatal = oldFatal
		DB = oldDB
	}()

	os.Setenv("DATABASE_URL", "postgres://test")

	newPool = func(ctx context.Context, url string) (*pgxpool.Pool, error) {

		if url != "postgres://test" {
			t.Fatalf("unexpected url: %s", url)
		}

		return &pgxpool.Pool{}, nil
	}

	ping = func(pool *pgxpool.Pool) error {
		return nil
	}

	sleep = func(time.Duration) {}

	fatal = func(v ...any) {
		t.Fatal("fatal should not be called")
	}

	Init()

	if DB == nil {
		t.Fatal("DB not initialized")
	}
}

func TestInit_RetryThenSuccess(t *testing.T) {

	oldNewPool := newPool
	oldPing := ping
	oldSleep := sleep
	oldFatal := fatal
	oldDB := DB

	defer func() {
		newPool = oldNewPool
		ping = oldPing
		sleep = oldSleep
		fatal = oldFatal
		DB = oldDB
	}()

	attempts := 0

	newPool = func(ctx context.Context, url string) (*pgxpool.Pool, error) {

		attempts++

		return &pgxpool.Pool{}, nil
	}

	ping = func(pool *pgxpool.Pool) error {

		if attempts < 3 {
			return errors.New("not ready")
		}

		return nil
	}

	sleep = func(time.Duration) {}

	fatal = func(v ...any) {
		t.Fatal("fatal should not be called")
	}

	Init()

	if attempts != 3 {
		t.Fatalf("expected 3 attempts got %d", attempts)
	}
}

func TestInit_NewPoolError(t *testing.T) {

	oldNewPool := newPool
	oldSleep := sleep
	oldFatal := fatal

	defer func() {
		newPool = oldNewPool
		sleep = oldSleep
		fatal = oldFatal
	}()

	called := false

	newPool = func(ctx context.Context, url string) (*pgxpool.Pool, error) {
		return nil, errors.New("connection failed")
	}

	sleep = func(time.Duration) {}

	fatal = func(v ...any) {
		called = true
	}

	Init()

	if !called {
		t.Fatal("expected fatal")
	}
}

func TestInit_DefaultURL(t *testing.T) {

	oldNewPool := newPool
	oldPing := ping
	oldSleep := sleep
	oldFatal := fatal
	oldDB := DB

	defer func() {
		newPool = oldNewPool
		ping = oldPing
		sleep = oldSleep
		fatal = oldFatal
		DB = oldDB
	}()

	os.Unsetenv("DATABASE_URL")

	newPool = func(ctx context.Context, url string) (*pgxpool.Pool, error) {

		expected := "postgres://postgres:postgres@127.0.0.1:5433/ai_orchestrator?sslmode=disable"

		if url != expected {
			t.Fatalf("expected default url got %s", url)
		}

		return &pgxpool.Pool{}, nil
	}

	ping = func(pool *pgxpool.Pool) error {
		return nil
	}

	sleep = func(time.Duration) {}

	fatal = func(v ...any) {
		t.Fatal("fatal should not be called")
	}

	Init()
}
