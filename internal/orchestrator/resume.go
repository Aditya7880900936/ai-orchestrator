package orchestrator

import (
	"encoding/json"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/repository"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

type analysisRepository interface {
	Save(resume string, result string) error
}

var (
	newResumeWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewResumeWorkflow()
	}

	executeResumePipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.ResumeAnalyzeResponse, error) = ExecutePipeline[models.ResumeAnalyzeResponse]

	newAnalysisRepository func() analysisRepository = func() analysisRepository {
		return repository.NewAnalysisRepository()
	}
)

func AnalyzeResume(req models.ResumeAnalyzeRequest) (*models.ResumeAnalyzeResponse, error) {

	cacheKey := cache.GenerateKey(req.ResumeText)

	resp, err := executeResumePipeline(
		cacheKey,
		req.ResumeText,
		newResumeWorkflow(),
	)
	if err != nil {
		return nil, err
	}

	repo := newAnalysisRepository()

	bytes, _ := json.Marshal(resp)

	if err := repo.Save(req.ResumeText, string(bytes)); err != nil {
		return nil, err
	}

	return resp, nil
}
