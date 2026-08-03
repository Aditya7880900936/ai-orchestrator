package retry

import (
	"errors"
	"testing"
	"time"
)

func TestExecute_SuccessFirstAttempt(t *testing.T) {

	calls := 0

	result, err := Execute[string](
		3,
		0*time.Millisecond,
		func() (string, error) {
			calls++
			return "success", nil
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "success" {
		t.Fatalf("expected success got %q", result)
	}

	if calls != 1 {
		t.Fatalf("expected 1 call got %d", calls)
	}
}

func TestExecute_SuccessAfterRetry(t *testing.T) {

	calls := 0

	result, err := Execute[string](
		3,
		0*time.Millisecond,
		func() (string, error) {

			calls++

			if calls < 3 {
				return "", errors.New("temporary failure")
			}

			return "success", nil
		},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if result != "success" {
		t.Fatalf("expected success got %q", result)
	}

	if calls != 3 {
		t.Fatalf("expected 3 calls got %d", calls)
	}
}

func TestExecute_FailureAfterMaxAttempts(t *testing.T) {

	calls := 0

	_, err := Execute[string](
		3,
		0,
		func() (string, error) {
			calls++
			return "", errors.New("always fails")
		},
	)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "always fails" {
		t.Fatalf("unexpected error: %v", err)
	}

	if calls != 3 {
		t.Fatalf("expected 3 calls got %d", calls)
	}
}

func TestExecute_ZeroAttempts(t *testing.T) {

	calls := 0

	result, err := Execute[string](
		0,
		0,
		func() (string, error) {
			calls++
			return "success", nil
		},
	)

	if calls != 0 {
		t.Fatalf("expected 0 calls got %d", calls)
	}

	if result != "" {
		t.Fatalf("expected zero value got %q", result)
	}

	if err != nil {
		t.Fatalf("expected nil error got %v", err)
	}
}
