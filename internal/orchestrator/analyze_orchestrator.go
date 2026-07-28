package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func Analyze(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

	cacheKey := cache.GenerateKey(req.Prompt)

	return ExecutePipeline[models.AnalyzeResponse](
		cacheKey,
		req.Prompt,
		workflow.NewAnalyzeWorkflow(),
	)
}