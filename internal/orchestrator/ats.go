package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func CalculateATS(req models.ATSScoreRequest) (*models.ATSScoreResponse, error) {
	cacheKey := cache.GenerateKey(req.ResumeText)

	return ExecutePipeline[models.ATSScoreResponse](
		cacheKey,
		req.ResumeText,
		workflow.NewATSWorkflow(),
	)
}