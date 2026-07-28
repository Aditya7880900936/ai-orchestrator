package model

type JobMatchRequest struct {
	ResumeText     string `json:"resume_text" binding:"required"`
	JobDescription string `json:"job_description" binding:"required"`
}

type JobMatchResponse struct {
	MatchPercentage int      `json:"match_percentage"`
	MatchedSkills   []string `json:"matched_skills"`
	MissingSkills   []string `json:"missing_skills"`
	Strengths       []string `json:"strengths"`
	Weaknesses      []string `json:"weaknesses"`
	Recommendations []string `json:"recommendations"`
}
