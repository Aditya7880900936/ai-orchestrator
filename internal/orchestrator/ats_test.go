package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestCalculateATS_Success(t *testing.T) {

	oldWorkflow := newATSWorkflow
	oldPipeline := executeATSPipeline

	defer func() {
		newATSWorkflow = oldWorkflow
		executeATSPipeline = oldPipeline
	}()

	newATSWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeATSPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ATSScoreResponse, error) {

		called = true

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "Backend Developer" {
			t.Fatalf("unexpected input %s", input)
		}

		if wf == nil {
			t.Fatal("workflow is nil")
		}
		return &models.ATSScoreResponse{
			OverallScore: 90,
			SectionScores: models.SectionScore{
				Contact:    10,
				Summary:    20,
				Skills:     20,
				Experience: 20,
				Education:  20,
			},
			MissingKeywords: []string{"Docker"},
			Strengths:       []string{"Go"},
			Weaknesses:      []string{"Kubernetes"},
			Suggestions:     []string{"Add projects"},
		}, nil
	}

	req := models.ATSScoreRequest{
		ResumeText: "Backend Developer",
	}

	resp, err := CalculateATS(req)

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("pipeline not called")
	}

	if resp.OverallScore != 90 {
		t.Fatal("wrong score")
	}

	if resp.SectionScores.Contact != 10 {
		t.Fatal("wrong contact score")
	}

	if len(resp.MissingKeywords) != 1 {
		t.Fatal("wrong keywords")
	}
}

func TestCalculateATS_Error(t *testing.T) {

	oldWorkflow := newATSWorkflow
	oldPipeline := executeATSPipeline

	defer func() {
		newATSWorkflow = oldWorkflow
		executeATSPipeline = oldPipeline
	}()

	newATSWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeATSPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ATSScoreResponse, error) {

		return nil, errors.New("pipeline failed")
	}

	resp, err := CalculateATS(models.ATSScoreRequest{
		ResumeText: "Backend Developer",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "pipeline failed" {
		t.Fatal(err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}
