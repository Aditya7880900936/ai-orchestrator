package model

type SkillExtractionRequest struct {
	ResumeText string `json:"resume_text" binding:"required"`
}

type SkillExtractionResponse struct {
	TechnicalSkills []string `json:"technical_skills"`
	Frameworks      []string `json:"frameworks"`
	Databases       []string `json:"databases"`
	Cloud           []string `json:"cloud"`
	Tools           []string `json:"tools"`
	SoftSkills      []string `json:"soft_skills"`
}
