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

func TestMatchJob_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/job-match",
		bytes.NewBufferString(`{invalid json}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	MatchJob(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestMatchJob_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/job-match",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	MatchJob(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestMatchJob_Error(t *testing.T) {

	setupHandlerTest(t)

	old := matchJob
	defer func() {
		matchJob = old
	}()

	matchJob = func(req models.JobMatchRequest) (*models.JobMatchResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		if req.JobDescription != "backend role" {
			t.Fatal("unexpected job description")
		}

		return nil, errors.New("matching failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/job-match",
		bytes.NewBufferString(`{
			"resume_text":"resume",
			"job_description":"backend role"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	MatchJob(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "matching failed") {
		t.Fatal("expected error message")
	}
}

func TestMatchJob_Success(t *testing.T) {

	setupHandlerTest(t)

	old := matchJob
	defer func() {
		matchJob = old
	}()

	matchJob = func(req models.JobMatchRequest) (*models.JobMatchResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		if req.JobDescription != "backend role" {
			t.Fatal("unexpected job description")
		}

		return &models.JobMatchResponse{
			MatchPercentage: 90,
			MatchedSkills:   []string{"Go", "Docker"},
			MissingSkills:   []string{"Kubernetes"},
			Strengths:       []string{"Backend"},
			Weaknesses:      []string{"AWS"},
			Recommendations: []string{"Learn Kubernetes"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/job-match",
		bytes.NewBufferString(`{
			"resume_text":"resume",
			"job_description":"backend role"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	MatchJob(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	resp := w.Body.String()

	if !strings.Contains(resp, `"match_percentage":90`) {
		t.Fatal("expected match percentage")
	}

	if !strings.Contains(resp, "Docker") {
		t.Fatal("expected matched skills")
	}

	if !strings.Contains(resp, "Kubernetes") {
		t.Fatal("expected missing skills")
	}

	if !strings.Contains(resp, "Backend") {
		t.Fatal("expected strengths")
	}

	if !strings.Contains(resp, "AWS") {
		t.Fatal("expected weaknesses")
	}

	if !strings.Contains(resp, "Learn Kubernetes") {
		t.Fatal("expected recommendations")
	}
}
