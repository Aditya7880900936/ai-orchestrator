package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newAnalyzeWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewAnalyzeWorkflow()
	}

	executeAnalyzePipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.AnalyzeResponse, error) = ExecutePipeline[models.AnalyzeResponse]
)

func Analyze(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

	cacheKey := cache.GenerateKey(req.Prompt)

	return executeAnalyzePipeline(
		cacheKey,
		req.Prompt,
		newAnalyzeWorkflow(),
	)
}
