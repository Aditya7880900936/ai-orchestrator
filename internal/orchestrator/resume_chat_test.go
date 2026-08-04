package orchestrator

import (
	"errors"
	"strings"
	"testing"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func TestChatWithResume_Success(t *testing.T) {

	oldGetSession := getSession
	oldGetConversation := getConversation
	oldSaveConversation := saveConversation
	oldWorkflow := newResumeChatWorkflow
	oldPipeline := executeResumeChatPipeline

	defer func() {
		getSession = oldGetSession
		getConversation = oldGetConversation
		saveConversation = oldSaveConversation
		newResumeChatWorkflow = oldWorkflow
		executeResumeChatPipeline = oldPipeline
	}()

	getSession = func(sessionID string) (string, error) {
		if sessionID != "123" {
			t.Fatal("wrong session id")
		}
		return "Go Backend Resume", nil
	}

	getConversation = func(sessionID string) (string, error) {
		return "Previous Chat", nil
	}

	saved := false

	saveConversation = func(sessionID, conversation string) error {

		saved = true

		if !strings.Contains(conversation, "User: Explain Redis") {
			t.Fatal("user message missing")
		}

		if !strings.Contains(conversation, "Assistant: Redis is a cache") {
			t.Fatal("assistant message missing")
		}

		return nil
	}

	newResumeChatWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumeChatPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeChatResponse, error) {

		if cacheKey == "" {
			t.Fatal("expected cache key")
		}

		if !strings.Contains(input, "Go Backend Resume") {
			t.Fatal("resume missing")
		}

		if !strings.Contains(input, "Previous Chat") {
			t.Fatal("conversation missing")
		}

		if !strings.Contains(input, "Explain Redis") {
			t.Fatal("question missing")
		}

		if wf == nil {
			t.Fatal("workflow nil")
		}

		return &models.ResumeChatResponse{
			Answer: "Redis is a cache",
		}, nil
	}

	resp, err := ChatWithResume(models.ResumeChatRequest{
		SessionID: "123",
		Question:  "Explain Redis",
	})

	if err != nil {
		t.Fatal(err)
	}

	if !saved {
		t.Fatal("conversation not saved")
	}

	if resp.Answer != "Redis is a cache" {
		t.Fatal("wrong answer")
	}
}

func TestChatWithResume_InvalidSession(t *testing.T) {

	old := getSession
	defer func() {
		getSession = old
	}()

	getSession = func(sessionID string) (string, error) {
		return "", errors.New("not found")
	}

	resp, err := ChatWithResume(models.ResumeChatRequest{
		SessionID: "bad",
		Question:  "Hello",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}

func TestChatWithResume_PipelineError(t *testing.T) {

	oldGetSession := getSession
	oldGetConversation := getConversation
	oldWorkflow := newResumeChatWorkflow
	oldPipeline := executeResumeChatPipeline

	defer func() {
		getSession = oldGetSession
		getConversation = oldGetConversation
		newResumeChatWorkflow = oldWorkflow
		executeResumeChatPipeline = oldPipeline
	}()

	getSession = func(string) (string, error) {
		return "Resume", nil
	}

	getConversation = func(string) (string, error) {
		return "", nil
	}

	newResumeChatWorkflow = func() workflow.Workflow {
		return &MockWorkflow{}
	}

	executeResumeChatPipeline = func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeChatResponse, error) {

		return nil, errors.New("pipeline failed")
	}

	resp, err := ChatWithResume(models.ResumeChatRequest{
		SessionID: "123",
		Question:  "Question",
	})

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "pipeline failed" {
		t.Fatal(err)
	}

	if resp != nil {
		t.Fatal("expected nil response")
	}
}
