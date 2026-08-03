package cache

import (
	"errors"
	"testing"
	"time"
)

func TestGet_Success(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {

		if key != "my-key" {
			t.Fatalf("expected my-key got %s", key)
		}

		return "cached-value", nil
	}

	value, err := Get("my-key")

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if value != "cached-value" {
		t.Fatalf("expected cached-value got %s", value)
	}
}

func TestGet_Error(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {
		return "", errors.New("redis down")
	}

	_, err := Get("my-key")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSet_Success(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {

		if key != "my-key" {
			t.Fatalf("wrong key")
		}

		if value != "my-value" {
			t.Fatalf("wrong value")
		}

		if ttl != 30*time.Minute {
			t.Fatalf("wrong ttl")
		}

		return nil
	}

	err := Set("my-key", "my-value")

	if err != nil {
		t.Fatalf("unexpected error")
	}
}

func TestSet_Error(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {
		return errors.New("redis error")
	}

	err := Set("my-key", "my-value")

	if err == nil {
		t.Fatal("expected error")
	}
}
