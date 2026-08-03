package orchestrator

import (
	"testing"
	"time"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestCalculateATS(t *testing.T) {

	// Backup original dependencies
	oldWorkflow := newATSWorkflow
	oldCacheGet := cacheGet
	oldCacheSet := cacheSet
	oldExecuteRetry := executeRetry
	oldExtractJSON := extractJSON

	defer func() {
		newATSWorkflow = oldWorkflow
		cacheGet = oldCacheGet
		cacheSet = oldCacheSet
		executeRetry = oldExecuteRetry
		extractJSON = oldExtractJSON
	}()

	// Cache miss
	cacheGet = func(key string) (string, error) {
		return "", nil
	}

	// Ignore cache write
	cacheSet = func(key, value string) error {
		return nil
	}

	// Mock workflow
	newATSWorkflow = func() workflow.Workflow {
		return &MockWorkflow{
			Response: `{
				"overall_score":90,
				"section_scores":{
					"contact":10,
					"summary":20,
					"skills":20,
					"experience":20,
					"education":20
				},
				"missing_keywords":["Docker"],
				"strengths":["Go"],
				"weaknesses":["Kubernetes"],
				"suggestions":["Add projects"]
			}`,
		}
	}

	// Execute workflow directly
	executeRetry = func(
		attempts int,
		delay time.Duration,
		fn func() (string, error),
	) (string, error) {
		return fn()
	}

	extractJSON = func(s string) string {
		return s
	}

	req := models.ATSScoreRequest{
		ResumeText: "Backend Developer",
	}

	resp, err := CalculateATS(req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.OverallScore != 90 {
		t.Fatalf("expected overall score 90, got %d", resp.OverallScore)
	}

	if len(resp.MissingKeywords) != 1 {
		t.Fatal("expected one missing keyword")
	}

	if resp.SectionScores.Contact != 10 {
		t.Fatal("section score mismatch")
	}
}
