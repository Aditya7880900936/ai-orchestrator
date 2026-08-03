package handler

import (
	"bytes"
	"errors"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/gin-gonic/gin"
)

func TestExtractSkills_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/skills",
		bytes.NewBufferString(`{invalid json}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ExtractSkills(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestExtractSkills_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/skills",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ExtractSkills(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestExtractSkills_Error(t *testing.T) {

	setupHandlerTest(t)

	old := extractSkills
	defer func() {
		extractSkills = old
	}()

	extractSkills = func(req model.SkillExtractionRequest) (*model.SkillExtractionResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		return nil, errors.New("skill extraction failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/skills",
		bytes.NewBufferString(`{
			"resume_text":"resume"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ExtractSkills(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "skill extraction failed") {
		t.Fatal("expected error message")
	}
}

func TestExtractSkills_Success(t *testing.T) {

	setupHandlerTest(t)

	old := extractSkills
	defer func() {
		extractSkills = old
	}()

	extractSkills = func(req model.SkillExtractionRequest) (*model.SkillExtractionResponse, error) {

		if req.ResumeText != "resume" {
			t.Fatal("unexpected resume")
		}

		return &model.SkillExtractionResponse{
			TechnicalSkills: []string{"Go", "C++"},
			Frameworks:      []string{"Gin", "Fiber"},
			Databases:       []string{"PostgreSQL", "Redis"},
			Cloud:           []string{"AWS"},
			Tools:           []string{"Docker", "Git"},
			SoftSkills:      []string{"Leadership"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/skills",
		bytes.NewBufferString(`{
			"resume_text":"resume"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ExtractSkills(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	resp := w.Body.String()

	if !strings.Contains(resp, "Go") {
		t.Fatal("expected technical skills")
	}

	if !strings.Contains(resp, "Gin") {
		t.Fatal("expected frameworks")
	}

	if !strings.Contains(resp, "PostgreSQL") {
		t.Fatal("expected databases")
	}

	if !strings.Contains(resp, "AWS") {
		t.Fatal("expected cloud")
	}

	if !strings.Contains(resp, "Docker") {
		t.Fatal("expected tools")
	}

	if !strings.Contains(resp, "Leadership") {
		t.Fatal("expected soft skills")
	}
}
