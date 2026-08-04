package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewResumeImproveWorkflow(t *testing.T) {

	w := NewResumeImproveWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestResumeImproveWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateResumeImproveLLM

	defer func() {
		generateResumeImproveLLM = oldGenerate
	}()

	called := false

	generateResumeImproveLLM = func(prompt string) (string, error) {

		called = true

		// Verify important prompt sections
		checks := []string{
			"Senior Technical Recruiter",
			"STRICT RULES",
			`"improved_summary"`,
			`"improved_experience"`,
			`"improved_projects"`,
			`"missing_sections"`,
			`"action_verbs"`,
			`"overall_suggestions"`,
			"Resume:",
			"Backend Resume",
		}

		for _, c := range checks {
			if !strings.Contains(prompt, c) {
				t.Fatalf("prompt missing: %s", c)
			}
		}

		return `{"improved_summary":"Improved Resume"}`, nil
	}

	w := NewResumeImproveWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"improved_summary":"Improved Resume"}` {
		t.Fatal("unexpected response")
	}
}

func TestResumeImproveWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateResumeImproveLLM

	defer func() {
		generateResumeImproveLLM = oldGenerate
	}()

	generateResumeImproveLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewResumeImproveWorkflow()

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
