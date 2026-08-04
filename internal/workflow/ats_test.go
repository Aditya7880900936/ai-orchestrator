package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewATSWorkflow(t *testing.T) {

	w := NewATSWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestATSWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateATSLLM

	defer func() {
		generateATSLLM = oldGenerate
	}()

	called := false

	generateATSLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "ATS resume evaluator") {
			t.Fatal("ATS instruction missing")
		}

		if !strings.Contains(prompt, `"overall_score"`) {
			t.Fatal("JSON schema missing")
		}

		if !strings.Contains(prompt, "Backend Resume") {
			t.Fatal("resume input missing")
		}

		return `{"overall_score":90}`, nil
	}

	w := NewATSWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"overall_score":90}` {
		t.Fatal("unexpected response")
	}
}

func TestATSWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateATSLLM

	defer func() {
		generateATSLLM = oldGenerate
	}()

	generateATSLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewATSWorkflow()

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
