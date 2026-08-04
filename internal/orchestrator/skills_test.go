package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestExtractSkills_Success(t *testing.T) {

	oldWorkflow := newSkillExtractionWorkflow
	oldPipeline := executeSkillPipeline

	defer func() {
		newSkillExtractionWorkflow = oldWorkflow
		executeSkillPipeline = oldPipeline
	}()

	newSkillExtractionWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	called := false

	executeSkillPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.SkillExtractionResponse, error) {

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

		return &models.SkillExtractionResponse{
			TechnicalSkills: []string{"Go", "C++"},
			Frameworks:      []string{"Gin", "Fiber"},
			Databases:       []string{"PostgreSQL", "Redis"},
			Cloud:           []string{"AWS"},
			Tools:           []string{"Docker", "Git"},
			SoftSkills:      []string{"Leadership"},
		}, nil
	}

	resp, err := ExtractSkills(models.SkillExtractionRequest{
		ResumeText: "resume text",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !called {
		t.Fatal("pipeline not called")
	}

	if len(resp.TechnicalSkills) != 2 {
		t.Fatal("expected 2 technical skills")
	}

	if resp.Frameworks[0] != "Gin" {
		t.Fatal("unexpected framework")
	}

	if resp.Databases[0] != "PostgreSQL" {
		t.Fatal("unexpected database")
	}

	if resp.Cloud[0] != "AWS" {
		t.Fatal("unexpected cloud")
	}

	if resp.Tools[0] != "Docker" {
		t.Fatal("unexpected tool")
	}

	if resp.SoftSkills[0] != "Leadership" {
		t.Fatal("unexpected soft skill")
	}
}

func TestExtractSkills_Error(t *testing.T) {

	oldWorkflow := newSkillExtractionWorkflow
	oldPipeline := executeSkillPipeline

	defer func() {
		newSkillExtractionWorkflow = oldWorkflow
		executeSkillPipeline = oldPipeline
	}()

	newSkillExtractionWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeSkillPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.SkillExtractionResponse, error) {

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

	resp, err := ExtractSkills(models.SkillExtractionRequest{
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
