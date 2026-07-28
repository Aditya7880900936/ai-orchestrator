package workflow

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

type ResumeChatWorkflow struct{}

func NewResumeChatWorkflow() *ResumeChatWorkflow {
	return &ResumeChatWorkflow{}
}

func (w *ResumeChatWorkflow) Run(input string) (string, error) {

	prompt := `
You are an expert Technical Recruiter and Senior Software Engineer.

You are given:

1. Candidate Resume
2. Previous Conversation
3. Current Question

Answer the CURRENT QUESTION while considering the previous conversation.

Rules:
- Use ONLY information from the resume.
- Use previous conversation only for context.
- Never invent projects, companies, skills, metrics, or achievements.
- If the answer is not present in the resume, clearly say so.
- Be concise, accurate, and professional.
- Return ONLY valid JSON.
- No markdown.
- No explanations outside JSON.

Return exactly:

{
  "answer": "string"
}

` + input

	return llm.Generate(prompt)
}
