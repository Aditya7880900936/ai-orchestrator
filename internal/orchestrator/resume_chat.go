package orchestrator

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func ChatWithResume(req models.ResumeChatRequest) (*models.ResumeChatResponse, error) {

	var resume string
	var err error

	// Store resume on first request
	if req.ResumeText != "" {
		resume = req.ResumeText

		if err := cache.SaveSession(req.SessionID, resume); err != nil {
			return nil, err
		}
	} else {
		resume, err = cache.GetSession(req.SessionID)
		if err != nil {
			return nil, fmt.Errorf("session not found")
		}
	}

	// Load previous conversation
	conversation, _ := cache.GetConversation(req.SessionID)

	input := fmt.Sprintf(
		`Resume:
%s

Conversation:
%s

Current Question:
%s`,
		resume,
		conversation,
		req.Question,
	)

	cacheKey := cache.GenerateKey(input)

	resp, err := ExecutePipeline[models.ResumeChatResponse](
		cacheKey,
		input,
		workflow.NewResumeChatWorkflow(),
	)
	if err != nil {
		return nil, err
	}

	// Append latest Q&A
	conversation += fmt.Sprintf(
		"\nUser: %s\nAssistant: %s\n",
		req.Question,
		resp.Answer,
	)

	_ = cache.SaveConversation(req.SessionID, conversation)

	return resp, nil
}