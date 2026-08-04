package workflow

import "github.com/Aditya7880900936/ai-orchestrator/internal/llm"

// Dependency injection point
var generateSkillLLM = llm.Generate

type SkillExtractionWorkflow struct{}

func NewSkillExtractionWorkflow() *SkillExtractionWorkflow {
	return &SkillExtractionWorkflow{}
}

func (w *SkillExtractionWorkflow) Run(resume string) (string, error) {

	prompt := `
You are an expert technical recruiter.

Extract skills from the following resume.

Return ONLY valid JSON.

{
  "technical_skills": [],
  "frameworks": [],
  "databases": [],
  "cloud": [],
  "tools": [],
  "soft_skills": []
}

Resume:
` + resume

	return generateSkillLLM(prompt)
}
