package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newSkillExtractionWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewSkillExtractionWorkflow()
	}

	executeSkillPipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.SkillExtractionResponse, error) = ExecutePipeline[models.SkillExtractionResponse]
)

func ExtractSkills(req models.SkillExtractionRequest) (*models.SkillExtractionResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	return executeSkillPipeline(
		cacheKey,
		req.ResumeText,
		newSkillExtractionWorkflow(),
	)
}
