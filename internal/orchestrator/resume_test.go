package orchestrator

import (
	"errors"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

type mockAnalysisRepository struct {
	save func(string, string) error
}

func (m *mockAnalysisRepository) Save(resume, result string) error {
	return m.save(resume, result)
}

func TestAnalyzeResume_Success(t *testing.T) {

	oldWorkflow := newResumeWorkflow
	oldPipeline := executeResumePipeline
	oldRepo := newAnalysisRepository

	defer func() {
		newResumeWorkflow = oldWorkflow
		executeResumePipeline = oldPipeline
		newAnalysisRepository = oldRepo
	}()

	newResumeWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeAnalyzeResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if input != "resume text" {
			t.Fatal("unexpected input")
		}

		if wf == nil {
			t.Fatal("workflow nil")
		}

		return &models.ResumeAnalyzeResponse{
			Summary:         "Backend Engineer",
			Skills:          []string{"Go", "Redis"},
			ExperienceYears: 2,
			Strengths:       []string{"Problem Solving"},
			MissingSkills:   []string{"Kubernetes"},
		}, nil
	}

	saved := false

	newAnalysisRepository = func() analysisRepository {
		return &mockAnalysisRepository{
			save: func(resume, result string) error {

				saved = true

				if resume != "resume text" {
					t.Fatal("wrong resume")
				}

				if result == "" {
					t.Fatal("empty json")
				}

				return nil
			},
		}
	}

	resp, err := AnalyzeResume(models.ResumeAnalyzeRequest{
		ResumeText: "resume text",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !saved {
		t.Fatal("repository save not called")
	}

	if resp.Summary != "Backend Engineer" {
		t.Fatal("wrong summary")
	}
}

func TestAnalyzeResume_PipelineError(t *testing.T) {

	oldWorkflow := newResumeWorkflow
	oldPipeline := executeResumePipeline

	defer func() {
		newResumeWorkflow = oldWorkflow
		executeResumePipeline = oldPipeline
	}()

	newResumeWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeAnalyzeResponse, error) {

		return nil, errors.New("pipeline failed")
	}

	resp, err := AnalyzeResume(models.ResumeAnalyzeRequest{
		ResumeText: "resume text",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestAnalyzeResume_SaveError(t *testing.T) {

	oldWorkflow := newResumeWorkflow
	oldPipeline := executeResumePipeline
	oldRepo := newAnalysisRepository

	defer func() {
		newResumeWorkflow = oldWorkflow
		executeResumePipeline = oldPipeline
		newAnalysisRepository = oldRepo
	}()

	newResumeWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumePipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeAnalyzeResponse, error) {

		return &models.ResumeAnalyzeResponse{
			Summary: "Backend Engineer",
		}, nil
	}

	newAnalysisRepository = func() analysisRepository {
		return &mockAnalysisRepository{
			save: func(resume, result string) error {
				return errors.New("db failed")
			},
		}
	}

	resp, err := AnalyzeResume(models.ResumeAnalyzeRequest{
		ResumeText: "resume text",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "db failed" {
		t.Fatal(err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}
