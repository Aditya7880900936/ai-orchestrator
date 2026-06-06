package orchestrator

import (
	"time"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser"
	"github.com/Aditya7880900936/ai-orchestrator/internal/retry"
	"github.com/Aditya7880900936/ai-orchestrator/internal/workflow"
)

func Analyze(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

	wf := workflow.NewAnalyzeWorkflow()

	resp, err := retry.Execute(
		3,
		2*time.Second,
		func() (string, error) {
			return wf.Run(req.Prompt)
		},
	)

	if err != nil {
		return nil, err
	}

	cleaned := parser.ExtractJSON(resp)

	parsed, err := parser.Parse[models.AnalyzeResponse](cleaned)
	if err != nil {
		return nil, err
	}

	return parsed, nil
}