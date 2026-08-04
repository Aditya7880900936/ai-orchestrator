package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestImproveResume_Success(t *testing.T) {

	oldWorkflow := newResumeImproveWorkflow
	oldPipeline := executeResumeImprovePipeline

	defer func() {
		newResumeImproveWorkflow = oldWorkflow
		executeResumeImprovePipeline = oldPipeline
	}()

	newResumeImproveWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeResumeImprovePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeImproveResponse, error) {

		called = true

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "resume text" {
			t.Fatalf("unexpected input: %s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return &models.ResumeImproveResponse{
			ImprovedSummary: "Improved backend engineer profile",
			ImprovedExperience: []string{
				"Built scalable APIs",
			},
			ImprovedProjects: []string{
				"AI Orchestrator",
			},
			MissingSections: []string{
				"Achievements",
			},
			ActionVerbs: []string{
				"Designed",
				"Implemented",
			},
			OverallSuggestions: []string{
				"Quantify impact",
			},
		}, nil
	}

	resp, err := ImproveResume(models.ResumeImproveRequest{
		ResumeText: "resume text",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("pipeline not called")
	}

	if resp.ImprovedSummary != "Improved backend engineer profile" {
		t.Fatal("wrong summary")
	}

	if len(resp.ImprovedExperience) != 1 {
		t.Fatal("wrong experience")
	}

	if len(resp.ImprovedProjects) != 1 {
		t.Fatal("wrong projects")
	}

	if resp.MissingSections[0] != "Achievements" {
		t.Fatal("wrong missing section")
	}

	if len(resp.ActionVerbs) != 2 {
		t.Fatal("wrong action verbs")
	}

	if resp.OverallSuggestions[0] != "Quantify impact" {
		t.Fatal("wrong suggestion")
	}
}

func TestImproveResume_Error(t *testing.T) {

	oldWorkflow := newResumeImproveWorkflow
	oldPipeline := executeResumeImprovePipeline

	defer func() {
		newResumeImproveWorkflow = oldWorkflow
		executeResumeImprovePipeline = oldPipeline
	}()

	newResumeImproveWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumeImprovePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeImproveResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "resume text" {
			t.Fatalf("unexpected input: %s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}

		return nil, errors.New("pipeline failed")
	}

	resp, err := ImproveResume(models.ResumeImproveRequest{
		ResumeText: "resume text",
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
