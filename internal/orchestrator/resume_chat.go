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

	// First request: save resume
	if req.ResumeText != "" {
		resume = req.ResumeText

		if err := cache.SaveSession(req.SessionID, resume); err != nil {
			return nil, err
		}
	} else {
		// Later requests: load resume
		resume, err = cache.GetSession(req.SessionID)
		if err != nil {
			return nil, fmt.Errorf("session not found")
		}
	}

	input := fmt.Sprintf(
		`Resume:
%s

Question:
%s`,
		resume,
		req.Question,
	)

	cacheKey := cache.GenerateKey(input)

	return ExecutePipeline[models.ResumeChatResponse](
		cacheKey,
		input,
		workflow.NewResumeChatWorkflow(),
	)
}