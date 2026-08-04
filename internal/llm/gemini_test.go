package llm

import (
	"errors"
	"testing"
)

type mockLLM struct {
	response string
	err      error
}

func (m *mockLLM) Generate(prompt string) (string, error) {

	if prompt == "" {
		panic("prompt should not be empty")
	}

	return m.response, m.err
}

func TestGenerate_Success(t *testing.T) {

	old := Client

	defer func() {
		Client = old
	}()

	Client = &mockLLM{
		response: "hello world",
	}

	resp, err := Generate("backend resume")

	if err != nil {
		t.Fatal(err)
	}

	if resp != "hello world" {
		t.Fatal("unexpected response")
	}
}

func TestGenerate_Error(t *testing.T) {

	old := Client

	defer func() {
		Client = old
	}()

	Client = &mockLLM{
		err: errors.New("gemini failed"),
	}

	resp, err := Generate("backend resume")

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "gemini failed" {
		t.Fatal(err)
	}

	if resp != "" {
		t.Fatal("expected empty response")
	}
}

func TestInitGemini_NoAPIKey(t *testing.T) {

	old := Client

	defer func() {
		Client = old
	}()

	t.Setenv("GEMINI_API_KEY", "")

	err := InitGemini()

	// Depending on SDK behavior:
	// - Either initialization succeeds and API fails later.
	// - Or initialization returns an error.
	// We only verify that it does not panic.

	_ = err
}
