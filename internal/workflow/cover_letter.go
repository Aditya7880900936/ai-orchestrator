package workflow

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

// Dependency injection point
var generateCoverLetterLLM = llm.Generate

type CoverLetterWorkflow struct{}

func NewCoverLetterWorkflow() *CoverLetterWorkflow {
	return &CoverLetterWorkflow{}
}

func (w *CoverLetterWorkflow) Run(input string) (string, error) {

	prompt := `
You are an experienced Technical Recruiter and Hiring Manager.

Your task is to write a professional, ATS-friendly, and personalized cover letter.

STRICT RULES
1. Return ONLY valid JSON.
2. Do NOT use Markdown.
3. Do NOT use headings, bullet points, or code fences.
4. Do NOT invent experience, achievements, or skills.
5. Use only the information provided.
6. Maintain a professional and enthusiastic tone.
7. The cover letter should be between 250 and 400 words.
8. Address the company and role naturally.
9. End with a professional closing.

Return exactly this JSON:

{
  "cover_letter": "string"
}

Candidate Information:

` + input

	return generateCoverLetterLLM(prompt)
}
