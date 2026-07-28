package orchestrator

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func ExtractSkills(req models.SkillExtractionRequest) (*models.SkillExtractionResponse, error) {
	cacheKey := cache.GenerateKey(req.ResumeText)

	return ExecutePipeline[models.SkillExtractionResponse](
		cacheKey,
		req.ResumeText,
		workflow.NewSkillExtractionWorkflow(),
	)
}
