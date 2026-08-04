package orchestrator

import (
	"fmt"
	"testing"
	"time"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
)

func TestExecutePipeline_CacheHit(t *testing.T) {

	// Backup original functions
	originalCacheGet := cacheGet
	originalExecuteRetry := executeRetry

	defer func() {
		cacheGet = originalCacheGet
		executeRetry = originalExecuteRetry
	}()

	// Mock cache hit
	cacheGet = func(key string) (string, error) {
		return `{
			"summary":"Backend Developer",
			"keywords":["Go","Redis"]
		}`, nil
	}

	// If cache works, workflow should NEVER execute
	executeRetry = func(
		attempts int,
		delay time.Duration,
		fn func() (string, error),
	) (string, error) {
		t.Fatal("workflow should not execute on cache hit")
		return "", nil
	}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"test-key",
		"dummy input",
		&MockWorkflow{},
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Summary != "Backend Developer" {
		t.Fatalf("expected 'Backend Developer', got '%s'", resp.Summary)
	}

	if len(resp.Keywords) != 2 {
		t.Fatalf("expected 2 keywords")
	}
}

func TestExecutePipeline_CacheMiss(t *testing.T) {

	// Backup original functions
	originalCacheGet := cacheGet
	originalCacheSet := cacheSet
	originalExecuteRetry := executeRetry
	originalExtractJSON := extractJSON

	defer func() {
		cacheGet = originalCacheGet
		cacheSet = originalCacheSet
		executeRetry = originalExecuteRetry
		extractJSON = originalExtractJSON
	}()

	// Cache miss
	cacheGet = func(key string) (string, error) {
		return "", fmt.Errorf("cache miss")
	}

	// Verify cache set is called
	cacheSaved := false
	cacheSet = func(key, value string) error {
		cacheSaved = true
		return nil
	}

	// Mock workflow execution
	executeRetry = func(
		attempts int,
		delay time.Duration,
		fn func() (string, error),
	) (string, error) {
		return fn()
	}

	// Don't modify JSON
	extractJSON = func(s string) string {
		return s
	}

	wf := &MockWorkflow{
		Response: `{
			"summary":"AI Engineer",
			"keywords":["Go","LLM"]
		}`,
	}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"test-key",
		"resume text",
		wf,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Summary != "AI Engineer" {
		t.Fatalf("expected AI Engineer, got %s", resp.Summary)
	}

	if !cacheSaved {
		t.Fatal("expected cache to be saved")
	}
}

func TestExecutePipeline_InvalidJSON(t *testing.T) {

	originalCacheGet := cacheGet
	originalExecuteRetry := executeRetry
	originalExtractJSON := extractJSON

	defer func() {
		cacheGet = originalCacheGet
		executeRetry = originalExecuteRetry
		extractJSON = originalExtractJSON
	}()

	cacheGet = func(key string) (string, error) {
		return "", fmt.Errorf("cache miss")
	}

	executeRetry = func(
		attempts int,
		delay time.Duration,
		fn func() (string, error),
	) (string, error) {
		return "this is not json", nil
	}

	extractJSON = func(s string) string {
		return s
	}

	wf := &MockWorkflow{}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"key",
		"input",
		wf,
	)

	if err == nil {
		t.Fatal("expected parsing error")
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestExecutePipeline_WorkflowError(t *testing.T) {

	// Backup
	originalCacheGet := cacheGet
	originalCacheSet := cacheSet
	originalExecuteRetry := executeRetry

	defer func() {
		cacheGet = originalCacheGet
		cacheSet = originalCacheSet
		executeRetry = originalExecuteRetry
	}()

	// Cache miss
	cacheGet = func(key string) (string, error) {
		return "", fmt.Errorf("cache miss")
	}

	// Cache should NEVER be called
	cacheSet = func(key, value string) error {
		t.Fatal("cache should not be written when workflow fails")
		return nil
	}

	// Workflow fails
	executeRetry = func(
		attempts int,
		delay time.Duration,
		fn func() (string, error),
	) (string, error) {
		return "", fmt.Errorf("workflow failed")
	}

	wf := &MockWorkflow{}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"test-key",
		"dummy input",
		wf,
	)

	if err == nil {
		t.Fatal("expected error")
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestExecutePipeline_NestedJSON(t *testing.T) {

	// Backup
	oldCacheGet := cacheGet
	oldCacheSet := cacheSet
	oldExecuteRetry := executeRetry
	oldExtractJSON := extractJSON

	defer func() {
		cacheGet = oldCacheGet
		cacheSet = oldCacheSet
		executeRetry = oldExecuteRetry
		extractJSON = oldExtractJSON
	}()

	// Cache miss
	cacheGet = func(key string) (string, error) {
		return "", fmt.Errorf("cache miss")
	}

	cacheSet = func(key, value string) error {
		return nil
	}

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

	wf := &MockWorkflow{
		Response: `{
			"answer":"{\"summary\":\"Nested Response\",\"keywords\":[\"Go\",\"Redis\"]}"
		}`,
	}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"nested-key",
		"dummy input",
		wf,
	)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Summary != "Nested Response" {
		t.Fatalf("expected Nested Response, got %s", resp.Summary)
	}

	if len(resp.Keywords) != 2 {
		t.Fatalf("expected 2 keywords")
	}
}

func TestExecutePipeline_InvalidCacheJSON(t *testing.T) {

	oldCacheGet := cacheGet
	oldCacheSet := cacheSet
	oldExecuteRetry := executeRetry
	oldExtractJSON := extractJSON

	defer func() {
		cacheGet = oldCacheGet
		cacheSet = oldCacheSet
		executeRetry = oldExecuteRetry
		extractJSON = oldExtractJSON
	}()

	cacheGet = func(string) (string, error) {
		return `invalid json`, nil
	}

	cacheSet = func(key, value string) error {
		return nil
	}

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

	wf := &MockWorkflow{
		Response: `{
			"summary":"Recovered",
			"keywords":["Go"]
		}`,
	}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"key",
		"input",
		wf,
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Summary != "Recovered" {
		t.Fatal("expected recovered response")
	}
}

func TestExecutePipeline_WrappedAnswerNotJSON(t *testing.T) {

	oldCacheGet := cacheGet
	oldExecuteRetry := executeRetry
	oldExtractJSON := extractJSON

	defer func() {
		cacheGet = oldCacheGet
		executeRetry = oldExecuteRetry
		extractJSON = oldExtractJSON
	}()

	cacheGet = func(string) (string, error) {
		return "", fmt.Errorf("miss")
	}

	executeRetry = func(
		a int,
		d time.Duration,
		fn func() (string, error),
	) (string, error) {
		return `{"answer":"hello world"}`, nil
	}

	extractJSON = func(s string) string {
		return s
	}

	_, err := ExecutePipeline[models.AnalyzeResponse](
		"key",
		"input",
		&MockWorkflow{},
	)

	if err == nil {
		t.Fatal("expected parse error")
	}
}

func TestExecutePipeline_CacheSetErrorIgnored(t *testing.T) {

	oldCacheGet := cacheGet
	oldCacheSet := cacheSet
	oldExecuteRetry := executeRetry
	oldExtractJSON := extractJSON

	defer func() {
		cacheGet = oldCacheGet
		cacheSet = oldCacheSet
		executeRetry = oldExecuteRetry
		extractJSON = oldExtractJSON
	}()

	cacheGet = func(string) (string, error) {
		return "", fmt.Errorf("miss")
	}

	cacheSet = func(key, value string) error {
		return fmt.Errorf("redis down")
	}

	executeRetry = func(
		a int,
		d time.Duration,
		fn func() (string, error),
	) (string, error) {
		return fn()
	}

	extractJSON = func(s string) string {
		return s
	}

	wf := &MockWorkflow{
		Response: `{
			"summary":"Success",
			"keywords":["Go"]
		}`,
	}

	resp, err := ExecutePipeline[models.AnalyzeResponse](
		"key",
		"input",
		wf,
	)

	if err != nil {
		t.Fatal(err)
	}

	if resp.Summary != "Success" {
		t.Fatal("expected success")
	}
}
