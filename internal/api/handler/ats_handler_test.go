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

func TestCalculateATS_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{invalid json}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/ats",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CalculateATS(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCalculateATS_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/ats",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CalculateATS(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestCalculateATS_Error(t *testing.T) {

	setupHandlerTest(t)

	old := calculateATS
	defer func() {
		calculateATS = old
	}()

	calculateATS = func(req models.ATSScoreRequest) (*models.ATSScoreResponse, error) {
		return nil, errors.New("ats failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"resume_text":"my resume"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/ats",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CalculateATS(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "ats failed") {
		t.Fatal("expected error message")
	}
}

func TestCalculateATS_Success(t *testing.T) {

	setupHandlerTest(t)

	old := calculateATS
	defer func() {
		calculateATS = old
	}()

	calculateATS = func(req models.ATSScoreRequest) (*models.ATSScoreResponse, error) {

		if req.ResumeText != "my resume" {
			t.Fatalf("unexpected resume text")
		}

		return &models.ATSScoreResponse{
			OverallScore: 92,
			SectionScores: models.SectionScore{
				Skills: 95,
			},
			Strengths: []string{"Go"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"resume_text":"my resume"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/ats",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	CalculateATS(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	if !strings.Contains(w.Body.String(), `"overall_score":92`) {
		t.Fatal("expected overall score")
	}

	if !strings.Contains(w.Body.String(), "Go") {
		t.Fatal("expected strengths")
	}
}