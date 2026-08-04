package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewAnalyzeWorkflow(t *testing.T) {

	w := NewAnalyzeWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestAnalyzeWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateLLM

	defer func() {
		generateLLM = oldGenerate
	}()

	called := false

	generateLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "You are an intelligent analyzer") {
			t.Fatal("prompt missing instruction")
		}

		if !strings.Contains(prompt, "Backend Resume") {
			t.Fatal("input missing")
		}

		return `{"summary":"Backend","keywords":["Go"]}`, nil
	}

	w := NewAnalyzeWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"summary":"Backend","keywords":["Go"]}` {
		t.Fatal("unexpected response")
	}
}

func TestAnalyzeWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateLLM

	defer func() {
		generateLLM = oldGenerate
	}()

	generateLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewAnalyzeWorkflow()

	resp, err := w.Run("Resume")

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
