package orchestrator

import (
	"encoding/json"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/repository"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func AnalyzeResume(req models.ResumeAnalyzeRequest) (*models.ResumeAnalyzeResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	resp, err := ExecutePipeline[models.ResumeAnalyzeResponse](
		cacheKey,
		req.ResumeText,
		workflow.NewResumeWorkflow(),
	)
	if err != nil {
		return nil, err
	}

	repo := repository.NewAnalysisRepository()

	bytes, _ := json.Marshal(resp)

	if err := repo.Save(req.ResumeText, string(bytes)); err != nil {
		return nil, err
	}

	return resp, nil
}
