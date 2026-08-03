package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func MatchJob(req models.JobMatchRequest) (*models.JobMatchResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText + req.JobDescription)

	input := "Resume:\n" + req.ResumeText + "\n\nJob Description:\n" + req.JobDescription

	return ExecutePipeline[models.JobMatchResponse](
		cacheKey,
		input,
		workflow.NewJobMatchWorkflow(),
	)
}
