package parser

import (
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
)

func TestParse_ValidJSON(t *testing.T) {

	raw := `{
		"summary":"Backend Engineer",
		"keywords":["Go","Redis"]
	}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Summary != "Backend Engineer" {
		t.Fatalf("expected Backend Engineer got %s", resp.Summary)
	}

	if len(resp.Keywords) != 2 {
		t.Fatalf("expected 2 keywords")
	}
}

func TestParse_WrappedJSON(t *testing.T) {

	raw := `{
		"answer":"{\"summary\":\"Wrapped Response\",\"keywords\":[\"Go\",\"Docker\"]}"
	}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp.Summary != "Wrapped Response" {
		t.Fatalf("expected Wrapped Response got %s", resp.Summary)
	}

	if len(resp.Keywords) != 2 {
		t.Fatalf("expected 2 keywords")
	}
}

func TestParse_InvalidJSON(t *testing.T) {

	raw := `this is not json`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err == nil {
		t.Fatal("expected error")
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestParse_UnknownFieldsIgnored(t *testing.T) {

	raw := `{
		"foo":"bar"
	}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp == nil {
		t.Fatal("expected response")
	}

	if resp.Summary != "" {
		t.Fatal("expected empty summary")
	}

	if len(resp.Keywords) != 0 {
		t.Fatal("expected no keywords")
	}
}

func TestParse_AnswerIsNotJSON(t *testing.T) {

	raw := `{
		"answer":"hello world"
	}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "wrapped answer is not json" {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestParse_EmptyJSON(t *testing.T) {

	raw := `{}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp == nil {
		t.Fatal("expected response")
	}

	if resp.Summary != "" {
		t.Fatal("expected empty summary")
	}

	if len(resp.Keywords) != 0 {
		t.Fatal("expected no keywords")
	}
}

func TestParse_EmptyAnswer(t *testing.T) {

	raw := `{
		"answer":""
	}`

	resp, err := Parse[models.AnalyzeResponse](raw)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "wrapped answer is not json" {
		t.Fatalf("unexpected error: %v", err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestParse_NestedAnswerArray(t *testing.T) {

	raw := `{
		"answer":"{\"answer\":[\"Go\",\"Redis\"]}"
	}`

	type Response struct {
		Answer string `json:"answer"`
	}

	resp, err := Parse[Response](raw)

	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	expected := "Go\nRedis"

	if resp.Answer != expected {
		t.Fatalf("expected %q got %q", expected, resp.Answer)
	}
}
