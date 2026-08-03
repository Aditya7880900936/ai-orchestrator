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

func TestChatWithResume_InvalidJSON(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/chat",
		bytes.NewBufferString(`{invalid json}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ChatWithResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestChatWithResume_EmptyRequest(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/chat",
		bytes.NewBufferString(`{}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ChatWithResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}
}

func TestChatWithResume_Error(t *testing.T) {

	setupHandlerTest(t)

	old := chatWithResume
	defer func() {
		chatWithResume = old
	}()

	chatWithResume = func(req models.ResumeChatRequest) (*models.ResumeChatResponse, error) {

		if req.SessionID != "session-123" {
			t.Fatal("unexpected session id")
		}

		if req.Question != "Tell me about my skills" {
			t.Fatal("unexpected question")
		}

		return nil, errors.New("chat failed")
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/chat",
		bytes.NewBufferString(`{
			"session_id":"session-123",
			"question":"Tell me about my skills"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ChatWithResume(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf("expected %d got %d", http.StatusInternalServerError, w.Code)
	}

	if !strings.Contains(w.Body.String(), "chat failed") {
		t.Fatal("expected error message")
	}
}

func TestChatWithResume_Success(t *testing.T) {

	setupHandlerTest(t)

	old := chatWithResume
	defer func() {
		chatWithResume = old
	}()

	chatWithResume = func(req models.ResumeChatRequest) (*models.ResumeChatResponse, error) {

		if req.SessionID != "session-123" {
			t.Fatal("unexpected session id")
		}

		if req.Question != "Tell me about my skills" {
			t.Fatal("unexpected question")
		}

		return &models.ResumeChatResponse{
			Answer: "You have strong Go and Docker skills.",
		}, nil
	}

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/chat",
		bytes.NewBufferString(`{
			"session_id":"session-123",
			"question":"Tell me about my skills"
		}`),
	)

	req.Header.Set("Content-Type", "application/json")
	c.Request = req

	ChatWithResume(c)

	if w.Code != http.StatusOK {
		t.Fatalf("expected %d got %d", http.StatusOK, w.Code)
	}

	if !strings.Contains(w.Body.String(), "Go and Docker") {
		t.Fatal("expected answer")
	}
}