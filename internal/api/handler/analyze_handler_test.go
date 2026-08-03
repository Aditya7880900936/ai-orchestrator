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

func TestAnalyzeHandler_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	gin.SetMode(gin.TestMode)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{invalid json}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	c.Request = req

	AnalyzeHandler(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

func TestAnalyzeHandler_EmptyPrompt(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"prompt":""
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")

	c.Request = req

	AnalyzeHandler(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf(
			"expected %d got %d",
			http.StatusBadRequest,
			w.Code,
		)
	}
}

func TestAnalyzeHandler_Success(t *testing.T) {

	setupHandlerTest(t)

	old := analyze
	defer func() {
		analyze = old
	}()

	analyze = func(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

		if req.Prompt != "hello" {
			t.Fatalf("unexpected prompt: %s", req.Prompt)
		}

		return &models.AnalyzeResponse{
			Summary:  "Backend Engineer",
			Keywords: []string{"Go", "Redis"},
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"prompt":"hello"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeHandler(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	if !strings.Contains(w.Body.String(), "Backend Engineer") {
		t.Fatal("response body does not contain summary")
	}
}

func TestAnalyzeHandler_Error(t *testing.T) {

	setupHandlerTest(t)

	old := analyze
	defer func() {
		analyze = old
	}()

	analyze = func(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {
		return nil, errors.New("llm failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := bytes.NewBufferString(`{
		"prompt":"hello"
	}`)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/analyze",
		body,
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	AnalyzeHandler(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf(
			"expected %d got %d",
			http.StatusInternalServerError,
			w.Code,
		)
	}

	if !strings.Contains(w.Body.String(), "llm failed") {
		t.Fatal("expected error message in response")
	}
}
