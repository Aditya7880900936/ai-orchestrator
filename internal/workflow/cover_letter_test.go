package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewCoverLetterWorkflow(t *testing.T) {

	w := NewCoverLetterWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestCoverLetterWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateCoverLetterLLM

	defer func() {
		generateCoverLetterLLM = oldGenerate
	}()

	called := false

	generateCoverLetterLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "professional, ATS-friendly, and personalized cover letter") {
			t.Fatal("cover letter instruction missing")
		}

		if !strings.Contains(prompt, `"cover_letter": "string"`) {
			t.Fatal("json schema missing")
		}

		if !strings.Contains(prompt, "Candidate Information:") {
			t.Fatal("candidate section missing")
		}

		if !strings.Contains(prompt, "Go Backend Engineer") {
			t.Fatal("candidate input missing")
		}

		return `{"cover_letter":"Dear Hiring Manager..."}`, nil
	}

	w := NewCoverLetterWorkflow()

	resp, err := w.Run("Go Backend Engineer")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"cover_letter":"Dear Hiring Manager..."}` {
		t.Fatal("unexpected response")
	}
}

func TestCoverLetterWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateCoverLetterLLM

	defer func() {
		generateCoverLetterLLM = oldGenerate
	}()

	generateCoverLetterLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewCoverLetterWorkflow()

	resp, err := w.Run("Go Backend Engineer")

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
