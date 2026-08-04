package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newATSWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewATSWorkflow()
	}

	executeATSPipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ATSScoreResponse, error) = ExecutePipeline[models.ATSScoreResponse]
)

func CalculateATS(req models.ATSScoreRequest) (*models.ATSScoreResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	return executeATSPipeline(
		cacheKey,
		req.ResumeText,
		newATSWorkflow(),
	)
}
