package cache

import (
	"errors"
	"testing"
	"time"
)

func TestSaveSession(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {

		if key != "resume:123" {
			t.Fatalf("wrong key")
		}

		if value != "resume-data" {
			t.Fatalf("wrong value")
		}

		if ttl != sessionTTL {
			t.Fatalf("wrong ttl")
		}

		return nil
	}

	if err := SaveSession("123", "resume-data"); err != nil {
		t.Fatal(err)
	}
}

func TestGetSession(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {

		if key != "resume:123" {
			t.Fatalf("wrong key")
		}

		return "resume-data", nil
	}

	res, err := GetSession("123")

	if err != nil {
		t.Fatal(err)
	}

	if res != "resume-data" {
		t.Fatal("unexpected result")
	}
}

func TestSaveConversation(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {

		if key != "chat:123" {
			t.Fatalf("wrong key")
		}

		return nil
	}

	if err := SaveConversation("123", "hello"); err != nil {
		t.Fatal(err)
	}
}

func TestGetConversation_Success(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {
		return "hello", nil
	}

	res, err := GetConversation("123")

	if err != nil {
		t.Fatal(err)
	}

	if res != "hello" {
		t.Fatal("unexpected result")
	}
}

func TestGetConversation_Error(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {
		return "", errors.New("redis down")
	}

	res, err := GetConversation("123")

	if err != nil {
		t.Fatal("expected nil error")
	}

	if res != "" {
		t.Fatal("expected empty string")
	}
}

func TestSaveSession_Error(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {
		return errors.New("redis error")
	}

	if err := SaveSession("123", "resume"); err == nil {
		t.Fatal("expected error")
	}
}

func TestGetSession_Error(t *testing.T) {

	old := redisGet
	defer func() {
		redisGet = old
	}()

	redisGet = func(key string) (string, error) {
		return "", errors.New("redis down")
	}

	_, err := GetSession("123")

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestSaveConversation_Error(t *testing.T) {

	old := redisSet
	defer func() {
		redisSet = old
	}()

	redisSet = func(key, value string, ttl time.Duration) error {
		return errors.New("redis error")
	}

	if err := SaveConversation("123", "hello"); err == nil {
		t.Fatal("expected error")
	}
}
