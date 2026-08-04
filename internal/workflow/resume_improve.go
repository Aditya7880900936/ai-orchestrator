package workflow

import (
	"github.com/Aditya7880900936/ai-orchestrator/internal/llm"
)

// Dependency injection point
var generateResumeImproveLLM = llm.Generate

type ResumeImproveWorkflow struct{}

func NewResumeImproveWorkflow() *ResumeImproveWorkflow {
	return &ResumeImproveWorkflow{}
}

func (w *ResumeImproveWorkflow) Run(resumeText string) (string, error) {

	prompt := `
You are a Senior Technical Recruiter and Staff Software Engineer with extensive experience reviewing resumes for leading technology companies.

Your task is to improve the resume while preserving factual accuracy.

OBJECTIVE
- Rewrite the resume to be ATS-friendly.
- Improve clarity, grammar, and professionalism.
- Highlight the candidate's strengths.
- Make every sentence concise and impactful.

STRICT RULES
1. NEVER invent experience, projects, skills, certifications, achievements, companies, dates, or metrics.
2. If information is missing, mention it only in "missing_sections".
3. Rewrite existing content using stronger and more professional language.
4. Start every experience bullet with a powerful action verb.
5. Keep recommendations practical and concise.
6. Return ONLY valid JSON.
7. Do NOT include markdown.
8. Do NOT use **bold**, bullet symbols, numbered lists, headings, HTML, or code fences.
9. Every array element must be plain text.
10. Every value must be either a string or an array of strings.
11. Ensure the JSON can be parsed directly with a standard JSON parser.
12. Do not include any explanation before or after the JSON.

Return EXACTLY this JSON schema:

{
  "improved_summary": "string",
  "improved_experience": [
    "string"
  ],
  "improved_projects": [
    "string"
  ],
  "missing_sections": [
    "string"
  ],
  "action_verbs": [
    "string"
  ],
  "overall_suggestions": [
    "string"
  ]
}

FIELD REQUIREMENTS

improved_summary
- Write 3-5 professional sentences.
- Focus on technical strengths.
- Keep it ATS-friendly.
- Do not fabricate information.

improved_experience
- Rewrite every experience point.
- Improve readability and professionalism.
- Begin each point with a strong action verb.
- Focus on achievements whenever possible.
- Do not invent numbers or metrics.

improved_projects
- If projects are present, rewrite their descriptions.
- If projects are missing, suggest project ideas only.
- Never fabricate completed projects.

missing_sections
Include only genuinely missing resume sections such as:
- Contact Information
- Education
- Skills
- Projects
- Certifications
- Achievements
- Open Source Contributions
- Publications
- Languages

action_verbs
Return 15-20 powerful resume action verbs.

overall_suggestions
Return 5-10 concise suggestions that can significantly improve the resume.

Resume:

` + resumeText

	return generateResumeImproveLLM(prompt)
}
