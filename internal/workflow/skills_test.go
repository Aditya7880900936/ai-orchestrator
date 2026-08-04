package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewSkillExtractionWorkflow(t *testing.T) {

	w := NewSkillExtractionWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestSkillExtractionWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateSkillLLM

	defer func() {
		generateSkillLLM = oldGenerate
	}()

	called := false

	generateSkillLLM = func(prompt string) (string, error) {

		called = true

		if !strings.Contains(prompt, "Extract skills from the following resume") {
			t.Fatal("instruction missing")
		}

		if !strings.Contains(prompt, `"technical_skills"`) {
			t.Fatal("json schema missing")
		}

		if !strings.Contains(prompt, "Backend Resume") {
			t.Fatal("resume missing")
		}

		return `{"technical_skills":["Go"]}`, nil
	}

	w := NewSkillExtractionWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"technical_skills":["Go"]}` {
		t.Fatal("unexpected response")
	}
}

func TestSkillExtractionWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateSkillLLM

	defer func() {
		generateSkillLLM = oldGenerate
	}()

	generateSkillLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewSkillExtractionWorkflow()

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
