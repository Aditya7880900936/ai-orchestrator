package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	model "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/gin-gonic/gin"
)

func TestAnalyzeResume_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{invalid json}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestAnalyzeResume_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestAnalyzeResume_Error(t *testing.T) {

	setupHandlerTest(t)

	old := analyzeResume
	defer func() {
		analyzeResume = old
	}()

	analyzeResume = func(req model.ResumeAnalyzeRequest) (*model.ResumeAnalyzeResponse, error) {
		return nil, errors.New("analysis failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"resume_text":"my resume"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeResume(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "analysis failed") {
		t.Fatal("expected error message")
	}
}

func TestAnalyzeResume_Success(t *testing.T) {

	setupHandlerTest(t)

	old := analyzeResume
	defer func() {
		analyzeResume = old
	}()

	analyzeResume = func(req model.ResumeAnalyzeRequest) (*model.ResumeAnalyzeResponse, error) {

		if req.ResumeText != "my resume" {
			t.Fatalf("unexpected resume text")
		}

		return &model.ResumeAnalyzeResponse{
			Summary:         "Backend Engineer",
			Skills:          []string{"Go", "Docker"},
			ExperienceYears: 3,
			Strengths:       []string{"Backend"},
			MissingSkills:   []string{"Kubernetes"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"resume_text":"my resume"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeResume(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	resp := w.Body.String()

	if !strings.Contains(resp, "Backend Engineer") {
		t.Fatal("expected summary")
	}

	if !strings.Contains(resp, "Docker") {
		t.Fatal("expected skills")
	}

	if !strings.Contains(resp, "Backend") {
		t.Fatal("expected strengths")
	}

	if !strings.Contains(resp, "Kubernetes") {
		t.Fatal("expected missing skills")
	}
}
