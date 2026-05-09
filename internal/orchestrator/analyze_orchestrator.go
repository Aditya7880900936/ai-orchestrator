package orchestrator

import (
	"encoding/json"
	"fmt"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
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

	resp, err := llm.Generate(enrichedPrompt)
	if err != nil {
		return nil, err
	}

	var parsed models.AnalyzeResponse

	err = json.Unmarshal([]byte(resp), &parsed)
	if err != nil {
		return nil, err
	}

	return &parsed, nil
}