package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func ImproveResume(req models.ResumeImproveRequest) (*models.ResumeImproveResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	return ExecutePipeline[models.ResumeImproveResponse](
		cacheKey,
		req.ResumeText,
		workflow.NewResumeImproveWorkflow(),
	)
}