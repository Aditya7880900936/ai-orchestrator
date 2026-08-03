package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/gin-gonic/gin"
)

func TestImproveResume_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/improve",
		bytes.NewBufferString(`{invalid json}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ImproveResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestImproveResume_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/improve",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ImproveResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestImproveResume_Error(t *testing.T) {

	setupHandlerTest(t)

	old := improveResume
	defer func() {
		improveResume = old
	}()

	improveResume = func(req models.ResumeImproveRequest) (*models.ResumeImproveResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		return nil, errors.New("improvement failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/improve",
		bytes.NewBufferString(`{
			"resume_text":"resume"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ImproveResume(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "improvement failed") {
		t.Fatal("expected error message")
	}
}

func TestImproveResume_Success(t *testing.T) {

	setupHandlerTest(t)

	old := improveResume
	defer func() {
		improveResume = old
	}()

	improveResume = func(req models.ResumeImproveRequest) (*models.ResumeImproveResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		return &models.ResumeImproveResponse{
			ImprovedSummary:    "Improved backend engineer summary",
			ImprovedExperience: []string{"Built scalable Go APIs"},
			ImprovedProjects:   []string{"AI Orchestrator"},
			MissingSections:    []string{"Certifications"},
			ActionVerbs:        []string{"Designed", "Implemented"},
			OverallSuggestions: []string{"Quantify achievements"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/improve",
		bytes.NewBufferString(`{
			"resume_text":"resume"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ImproveResume(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	resp := w.Body.String()

	if !strings.Contains(resp, "Improved backend engineer summary") {
		t.Fatal("expected improved summary")
	}

	if !strings.Contains(resp, "Built scalable Go APIs") {
		t.Fatal("expected improved experience")
	}

	if !strings.Contains(resp, "AI Orchestrator") {
		t.Fatal("expected improved projects")
	}

	if !strings.Contains(resp, "Certifications") {
		t.Fatal("expected missing sections")
	}

	if !strings.Contains(resp, "Designed") {
		t.Fatal("expected action verbs")
	}

	if !strings.Contains(resp, "Quantify achievements") {
		t.Fatal("expected suggestions")
	}
}