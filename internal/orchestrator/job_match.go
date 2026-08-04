package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newJobMatchWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewJobMatchWorkflow()
	}

	executeJobMatchPipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.JobMatchResponse, error) = ExecutePipeline[models.JobMatchResponse]
)

func MatchJob(req models.JobMatchRequest) (*models.JobMatchResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText + req.JobDescription)

	input := "Resume:\n" +
		req.ResumeText +
		"\n\nJob Description:\n" +
		req.JobDescription

	return executeJobMatchPipeline(
		cacheKey,
		input,
		newJobMatchWorkflow(),
	)
}
