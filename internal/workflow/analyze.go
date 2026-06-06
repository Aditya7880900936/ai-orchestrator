package workflow

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

type AnalyzeWorkflow struct{}

func NewAnalyzeWorkflow() *AnalyzeWorkflow {
	return &AnalyzeWorkflow{}
}

func (w *AnalyzeWorkflow) Run(input string) (string, error) {

	prompt := fmt.Sprintf(`
You are an intelligent analyzer.

Return ONLY valid JSON:
{
  "summary": "...",
  "keywords": ["...", "..."]
}

Input:
%s
`, input)

	return llm.Generate(prompt)
}