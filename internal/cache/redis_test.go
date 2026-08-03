package cache

import (
	"os"
	"testing"
)

func TestInitRedis_DefaultAddress(t *testing.T) {

	old := os.Getenv("REDIS_ADDR")
	defer os.Setenv("REDIS_ADDR", old)

	os.Unsetenv("REDIS_ADDR")

	InitRedis()

	if Client == nil {
		t.Fatal("expected redis client")
	}
}

func TestInitRedis_CustomAddress(t *testing.T) {

	old := os.Getenv("REDIS_ADDR")
	defer os.Setenv("REDIS_ADDR", old)

	os.Setenv("REDIS_ADDR", "127.0.0.1:6380")

	InitRedis()

	if Client == nil {
		t.Fatal("expected redis client")
	}
}
