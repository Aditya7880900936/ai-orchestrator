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

func TestGenerateCoverLetter_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{invalid json}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/cover-letter",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	GenerateCoverLetter(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGenerateCoverLetter_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/cover-letter",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	GenerateCoverLetter(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestGenerateCoverLetter_Error(t *testing.T) {

	setupHandlerTest(t)

	old := generateCoverLetter
	defer func() {
		generateCoverLetter = old
	}()

	generateCoverLetter = func(req models.CoverLetterRequest) (*models.CoverLetterResponse, error) {

		if req.Name != "Aditya" {
			t.Fatal("unexpected name")
		}

		if req.Company != "Google" {
			t.Fatal("unexpected company")
		}

		if req.Position != "Backend Engineer" {
			t.Fatal("unexpected position")
		}

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		if req.JobDescription != "backend role" {
			t.Fatal("unexpected job description")
		}

		return nil, errors.New("generation failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"name":"Aditya",
		"company":"Google",
		"position":"Backend Engineer",
		"resume_text":"resume",
		"job_description":"backend role"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/cover-letter",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	GenerateCoverLetter(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "generation failed") {
		t.Fatal("expected error message")
	}
}

func TestGenerateCoverLetter_Success(t *testing.T) {

	setupHandlerTest(t)

	old := generateCoverLetter
	defer func() {
		generateCoverLetter = old
	}()

	generateCoverLetter = func(req models.CoverLetterRequest) (*models.CoverLetterResponse, error) {

		if req.Name != "Aditya" {
			t.Fatal("unexpected name")
		}

		if req.Company != "Google" {
			t.Fatal("unexpected company")
		}

		if req.Position != "Backend Engineer" {
			t.Fatal("unexpected position")
		}

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		if req.JobDescription != "backend role" {
			t.Fatal("unexpected job description")
		}

		return &models.CoverLetterResponse{
			CoverLetter: "Dear Hiring Manager,\nI am excited to apply...",
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"name":"Aditya",
		"company":"Google",
		"position":"Backend Engineer",
		"resume_text":"resume",
		"job_description":"backend role"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/cover-letter",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	GenerateCoverLetter(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	if !strings.Contains(w.Body.String(), "Dear Hiring Manager") {
		t.Fatal("expected cover letter")
	}
}