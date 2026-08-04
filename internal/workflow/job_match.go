package workflow

import (
	"fmt"

	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

// Dependency injection point
var generateJobMatchLLM = llm.Generate

type JobMatchWorkflow struct{}

func NewJobMatchWorkflow() *JobMatchWorkflow {
	return &JobMatchWorkflow{}
}

func (w *JobMatchWorkflow) Run(input string) (string, error) {

	prompt := fmt.Sprintf(`
You are an expert technical recruiter.

Compare the resume with the job description.

Return ONLY valid JSON.

{
  "match_percentage": 0,
  "matched_skills": [],
  "missing_skills": [],
  "strengths": [],
  "weaknesses": [],
  "recommendations": []
}

%s
`, input)

	return generateJobMatchLLM(prompt)
}
