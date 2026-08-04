package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newResumeImproveWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewResumeImproveWorkflow()
	}

	executeResumeImprovePipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeImproveResponse, error) = ExecutePipeline[models.ResumeImproveResponse]
)

func ImproveResume(req models.ResumeImproveRequest) (*models.ResumeImproveResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	return executeResumeImprovePipeline(
		cacheKey,
		req.ResumeText,
		newResumeImproveWorkflow(),
	)
}
