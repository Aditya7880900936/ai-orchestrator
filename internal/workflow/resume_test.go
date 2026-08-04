package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewResumeWorkflow(t *testing.T) {

	w := NewResumeWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestResumeWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateResumeLLM

	defer func() {
		generateResumeLLM = oldGenerate
	}()

	called := false

	generateResumeLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "expert technical recruiter") {
			t.Fatal("prompt missing recruiter instruction")
		}

		if !strings.Contains(prompt, `"experience_years"`) {
			t.Fatal("json schema missing")
		}

		if !strings.Contains(prompt, "Backend Resume") {
			t.Fatal("resume input missing")
		}

		return `{"summary":"Backend Engineer"}`, nil
	}

	w := NewResumeWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"summary":"Backend Engineer"}` {
		t.Fatal("unexpected response")
	}
}

func TestResumeWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateResumeLLM

	defer func() {
		generateResumeLLM = oldGenerate
	}()

	generateResumeLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewResumeWorkflow()

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
