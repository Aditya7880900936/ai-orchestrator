package cache

import "testing"

func TestGenerateKey_SameInput(t *testing.T) {

	input := "Backend Engineer"

	hash1 := GenerateKey(input)
	hash2 := GenerateKey(input)

	if hash1 != hash2 {
		t.Fatal("expected same hash")
	}
}

func TestGenerateKey_DifferentInput(t *testing.T) {

	hash1 := GenerateKey("Go")
	hash2 := GenerateKey("Redis")

	if hash1 == hash2 {
		t.Fatal("expected different hashes")
	}
}

func TestGenerateKey_Length(t *testing.T) {

	hash := GenerateKey("hello")

	if len(hash) != 64 {
		t.Fatalf("expected length 64 got %d", len(hash))
	}
}

func TestGenerateKey_EmptyString(t *testing.T) {

	hash := GenerateKey("")

	if hash == "" {
		t.Fatal("expected hash")
	}

	if len(hash) != 64 {
		t.Fatalf("expected length 64 got %d", len(hash))
	}
}
