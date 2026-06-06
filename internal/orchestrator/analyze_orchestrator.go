package orchestrator

import (
	"fmt"
	"time"

	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser"
	"github.com/Aditya7880900936/ai-orchestrator/internal/retry"
)

func Analyze(req models.AnalyzeRequest) (*models.AnalyzeResponse, error) {

	enrichedPrompt := fmt.Sprintf(`
You are an intelligent analyzer.

Return ONLY valid JSON in this format:
{
  "summary": "...",
  "keywords": ["...", "..."]
}

User Input:
%s
`, req.Prompt)

	resp, err := retry.Execute(
		3,
		2*time.Second,
		func() (string, error) {
			return llm.Generate(enrichedPrompt)
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