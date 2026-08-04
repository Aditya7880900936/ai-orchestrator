package workflow

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

// Dependency injection point
var generateResumeChatLLM = llm.Generate

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

Instructions:
- Answer ONLY the current question.
- Use ONLY the information present in the resume.
- Use previous conversation only as context.
- Do NOT invent any skills, companies, projects or achievements.
- If the answer is unavailable, say so clearly.

IMPORTANT:
Return ONLY a single JSON object.
Do NOT wrap JSON inside a string.
Do NOT return nested JSON.
Do NOT use markdown or code fences.

The response MUST exactly follow this schema:

{
  "answer": "your answer here"
}

Resume, conversation and question:

` + input

	return generateResumeChatLLM(prompt)
}
