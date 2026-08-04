package orchestrator

import (
	"errors"
	"strings"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestGenerateCoverLetter_Success(t *testing.T) {

	oldWorkflow := newCoverLetterWorkflow
	oldPipeline := executeCoverLetterPipeline

	defer func() {
		newCoverLetterWorkflow = oldWorkflow
		executeCoverLetterPipeline = oldPipeline
	}()

	newCoverLetterWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeCoverLetterPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.CoverLetterResponse, error) {

		called = true

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		// Verify formatted prompt
		if !strings.Contains(input, "Name: Aditya") {
			t.Fatal("name missing")
		}

		if !strings.Contains(input, "Company: OpenAI") {
			t.Fatal("company missing")
		}

		if !strings.Contains(input, "Position: Backend Engineer") {
			t.Fatal("position missing")
		}

		if !strings.Contains(input, "Resume:\nGo Developer") {
			t.Fatal("resume missing")
		}

		if !strings.Contains(input, "Job Description:\nBuild APIs") {
			t.Fatal("job description missing")
		}

		return &models.CoverLetterResponse{
			CoverLetter: "Dear Hiring Manager...",
		}, nil
	}

	resp, err := GenerateCoverLetter(models.CoverLetterRequest{
		Name:           "Aditya",
		Company:        "OpenAI",
		Position:       "Backend Engineer",
		ResumeText:     "Go Developer",
		JobDescription: "Build APIs",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("pipeline not called")
	}

	if resp.CoverLetter != "Dear Hiring Manager..." {
		t.Fatal("unexpected cover letter")
	}
}

func TestGenerateCoverLetter_Error(t *testing.T) {

	oldWorkflow := newCoverLetterWorkflow
	oldPipeline := executeCoverLetterPipeline

	defer func() {
		newCoverLetterWorkflow = oldWorkflow
		executeCoverLetterPipeline = oldPipeline
	}()

	newCoverLetterWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeCoverLetterPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.CoverLetterResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return nil, errors.New("pipeline failed")
	}

	resp, err := GenerateCoverLetter(models.CoverLetterRequest{
		Name:           "Aditya",
		Company:        "OpenAI",
		Position:       "Backend Engineer",
		ResumeText:     "Go Developer",
		JobDescription: "Build APIs",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "pipeline failed" {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}
