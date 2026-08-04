package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewJobMatchWorkflow(t *testing.T) {

	w := NewJobMatchWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestJobMatchWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateJobMatchLLM

	defer func() {
		generateJobMatchLLM = oldGenerate
	}()

	called := false

	generateJobMatchLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "Compare the resume with the job description") {
			t.Fatal("instruction missing")
		}

		if !strings.Contains(prompt, `"match_percentage"`) {
			t.Fatal("json schema missing")
		}

		if !strings.Contains(prompt, "Backend Resume") {
			t.Fatal("input missing")
		}

		return `{"match_percentage":90}`, nil
	}

	w := NewJobMatchWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"match_percentage":90}` {
		t.Fatal("unexpected response")
	}
}

func TestJobMatchWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateJobMatchLLM

	defer func() {
		generateJobMatchLLM = oldGenerate
	}()

	generateJobMatchLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewJobMatchWorkflow()

	resp, err := w.Run("Backend Resume")

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "llm failed" {
		t.Fatal(err)
	}

	if resp != "" {
		t.Fatal("expected empty response")
	}
}
