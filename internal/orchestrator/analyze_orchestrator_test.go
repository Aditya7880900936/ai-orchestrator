package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestAnalyze_Success(t *testing.T) {

	oldWorkflow := newAnalyzeWorkflow
	oldExecute := executeAnalyzePipeline

	defer func() {
		newAnalyzeWorkflow = oldWorkflow
		executeAnalyzePipeline = oldExecute
	}()

	newAnalyzeWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeAnalyzePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.AnalyzeResponse, error) {

		called = true

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "my resume" {
			t.Fatalf("unexpected input %s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return &models.AnalyzeResponse{
			Summary: "Backend Engineer",
			Keywords: []string{
				"Go",
				"Redis",
			},
		}, nil
	}

	resp, err := Analyze(models.AnalyzeRequest{
		Prompt: "my resume",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("ExecutePipeline not called")
	}

	if resp.Summary != "Backend Engineer" {
		t.Fatal("unexpected summary")
	}

	if len(resp.Keywords) != 2 {
		t.Fatal("expected 2 keywords")
	}
}

func TestAnalyze_Error(t *testing.T) {

	oldWorkflow := newAnalyzeWorkflow
	oldExecute := executeAnalyzePipeline

	defer func() {
		newAnalyzeWorkflow = oldWorkflow
		executeAnalyzePipeline = oldExecute
	}()

	newAnalyzeWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeAnalyzePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.AnalyzeResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "resume" {
			t.Fatalf("unexpected input %s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return nil, errors.New("pipeline failed")
	}

	resp, err := Analyze(models.AnalyzeRequest{
		Prompt: "resume",
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
