package workflow

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

type ATSWorkflow struct{}

func NewATSWorkflow() *ATSWorkflow {
	return &ATSWorkflow{}
}

func (w *ATSWorkflow) Run(input string) (string, error) {

	prompt := fmt.Sprintf(`
You are an ATS resume evaluator.

Analyze the following resume and return ONLY valid JSON.

{
  "overall_score": 0,
  "section_scores": {
    "contact": 0,
    "summary": 0,
    "skills": 0,
    "experience": 0,
    "education": 0
  },
  "missing_keywords": [],
  "strengths": [],
  "weaknesses": [],
  "suggestions": []
}

Resume:

%s
`, input)

	return llm.Generate(prompt)
}