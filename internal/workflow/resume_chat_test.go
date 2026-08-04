package workflow

import (
	"errors"
	"strings"
	"testing"
)

func TestNewResumeChatWorkflow(t *testing.T) {

	w := NewResumeChatWorkflow()

	if w == nil {
		t.Fatal("expected workflow")
	}
}

func TestResumeChatWorkflow_Run_Success(t *testing.T) {

	oldGenerate := generateResumeChatLLM

	defer func() {
		generateResumeChatLLM = oldGenerate
	}()

	called := false

	generateResumeChatLLM = func(prompt string) (string, error) {

		called = true

		checks := []string{
			"expert Technical Recruiter and Senior Software Engineer",
			"Candidate Resume",
			"Previous Conversation",
			"Current Question",
			`"answer": "your answer here"`,
			"Resume, conversation and question:",
			"Backend Resume",
		}

		for _, c := range checks {
			if !strings.Contains(prompt, c) {
				t.Fatalf("prompt missing: %s", c)
			}
		}

		return `{"answer":"Redis is an in-memory database"}`, nil
	}

	w := NewResumeChatWorkflow()

	resp, err := w.Run("Backend Resume")

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("Generate not called")
	}

	if resp != `{"answer":"Redis is an in-memory database"}` {
		t.Fatal("unexpected response")
	}
}

func TestResumeChatWorkflow_Run_Error(t *testing.T) {

	oldGenerate := generateResumeChatLLM

	defer func() {
		generateResumeChatLLM = oldGenerate
	}()

	generateResumeChatLLM = func(prompt string) (string, error) {
		return "", errors.New("llm failed")
	}

	w := NewResumeChatWorkflow()

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
