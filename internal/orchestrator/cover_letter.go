package orchestrator

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func GenerateCoverLetter(req models.CoverLetterRequest) (*models.CoverLetterResponse, error) {

	input := fmt.Sprintf(
		`Name: %s

Company: %s

Position: %s

Resume:
%s

Job Description:
%s`,
		req.Name,
		req.Company,
		req.Position,
		req.ResumeText,
		req.JobDescription,
	)

	cacheKey := cache.GenerateKey(input)

	return ExecutePipeline[models.CoverLetterResponse](
		cacheKey,
		input,
		workflow.NewCoverLetterWorkflow(),
	)
}