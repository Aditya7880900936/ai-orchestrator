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

Answer the user's question using ONLY the resume provided.

Rules:
1. Use only the information from the resume.
2. Do not invent projects, skills, companies, achievements, or metrics.
3. If the answer is not available in the resume, clearly say so.
4. Be concise and professional.
5. Return ONLY valid JSON.
6. Do NOT use Markdown.
7. Do NOT include explanations outside the JSON.

Return exactly:

{
  "answer": "string"
}

Resume and Question:

` + input

	return llm.Generate(prompt)
}