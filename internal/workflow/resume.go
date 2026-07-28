package workflow

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

type ResumeWorkflow struct{}

func NewResumeWorkflow() *ResumeWorkflow {
	return &ResumeWorkflow{}
}

func (w *ResumeWorkflow) Run(resumeText string) (string, error) {

	prompt := `
You are an expert technical recruiter.

Analyze the following resume and return ONLY valid JSON.

{
  "summary": "",
  "skills": [],
  "experience_years": 0,
  "strengths": [],
  "missing_skills": []
}

Resume:
` + resumeText

	return llm.Generate(prompt)
}