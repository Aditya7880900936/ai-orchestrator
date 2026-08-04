package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestMatchJob_Success(t *testing.T) {

	oldWorkflow := newJobMatchWorkflow
	oldPipeline := executeJobMatchPipeline

	defer func() {
		newJobMatchWorkflow = oldWorkflow
		executeJobMatchPipeline = oldPipeline
	}()

	newJobMatchWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeJobMatchPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.JobMatchResponse, error) {

		called = true

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		expectedInput := "Resume:\nresume text\n\nJob Description:\ngo backend job"

		if input != expectedInput {
			t.Fatalf("unexpected input:\n%s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return &models.JobMatchResponse{
			MatchPercentage: 90,
			MatchedSkills: []string{
				"Go",
				"Redis",
			},
			MissingSkills: []string{
				"Kubernetes",
			},
			Strengths: []string{
				"Backend",
			},
			Weaknesses: []string{
				"DevOps",
			},
			Recommendations: []string{
				"Learn Kubernetes",
			},
		}, nil
	}

	resp, err := MatchJob(models.JobMatchRequest{
		ResumeText:     "resume text",
		JobDescription: "go backend job",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("pipeline not called")
	}

	if resp.MatchPercentage != 90 {
		t.Fatal("wrong match percentage")
	}

	if len(resp.MatchedSkills) != 2 {
		t.Fatal("expected 2 matched skills")
	}

	if resp.MissingSkills[0] != "Kubernetes" {
		t.Fatal("wrong missing skill")
	}
}

func TestMatchJob_Error(t *testing.T) {

	oldWorkflow := newJobMatchWorkflow
	oldPipeline := executeJobMatchPipeline

	defer func() {
		newJobMatchWorkflow = oldWorkflow
		executeJobMatchPipeline = oldPipeline
	}()

	newJobMatchWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeJobMatchPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.JobMatchResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return nil, errors.New("pipeline failed")
	}

	resp, err := MatchJob(models.JobMatchRequest{
		ResumeText:     "resume text",
		JobDescription: "go backend job",
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
