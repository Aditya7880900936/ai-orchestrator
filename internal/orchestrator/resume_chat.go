package orchestrator

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	getSession       = cache.GetSession
	getConversation  = cache.GetConversation
	saveConversation = cache.SaveConversation

	newResumeChatWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewResumeChatWorkflow()
	}

	executeResumeChatPipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeChatResponse, error) = ExecutePipeline[models.ResumeChatResponse]
)

func ChatWithResume(req models.ResumeChatRequest) (*models.ResumeChatResponse, error) {

	resume, err := getSession(req.SessionID)
	if err != nil || resume == "" {
		return nil, fmt.Errorf("invalid session id or resume not found")
	}

	conversation, _ := getConversation(req.SessionID)

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

	resp, err := executeResumeChatPipeline(
		cacheKey,
		input,
		newResumeChatWorkflow(),
	)
	if err != nil {
		return nil, err
	}

	conversation += fmt.Sprintf(
		"\nUser: %s\nAssistant: %v\n",
		req.Question,
		resp.Answer,
	)

	_ = saveConversation(req.SessionID, conversation)

	return resp, nil
}
