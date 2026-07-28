package orchestrator

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func ChatWithResume(req models.ResumeChatRequest) (*models.ResumeChatResponse, error) {

	// Resume hamesha Redis se load hoga
	resume, err := cache.GetSession(req.SessionID)
	if err != nil || resume == "" {
		return nil, fmt.Errorf("invalid session id or resume not found")
	}

	// Previous conversation load karo
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

	// Conversation update
	conversation += fmt.Sprintf(
		"\nUser: %s\nAssistant: %s\n",
		req.Question,
		resp.Answer,
	)

	_ = cache.SaveConversation(req.SessionID, conversation)

	return resp, nil
}