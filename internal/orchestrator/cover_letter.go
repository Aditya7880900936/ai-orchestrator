package orchestrator

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

// Dependency injection points
var (
	newCoverLetterWorkflow func() workflow.Workflow = func() workflow.Workflow {
		return workflow.NewCoverLetterWorkflow()
	}

	executeCoverLetterPipeline func(
		cacheKey string,
		input string,
		wf workflow.Workflow,
	) (*models.CoverLetterResponse, error) = ExecutePipeline[models.CoverLetterResponse]
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

	return executeCoverLetterPipeline(
		cacheKey,
		input,
		newCoverLetterWorkflow(),
	)
}
