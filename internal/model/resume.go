package model

type ResumeAnalyzeRequest struct {
	ResumeText string `json:"resume_text" binding:"required"`
}

type ResumeAnalyzeResponse struct {
	Summary         string   `json:"summary"`
	Skills          []string `json:"skills"`
	ExperienceYears int      `json:"experience_years"`
	Strengths       []string `json:"strengths"`
	MissingSkills   []string `json:"missing_skills"`
}
