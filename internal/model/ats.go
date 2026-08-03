package model

type ATSScoreRequest struct {
	ResumeText string `json:"resume_text" binding:"required"`
}

type SectionScore struct {
	Contact    int `json:"contact"`
	Summary    int `json:"summary"`
	Skills     int `json:"skills"`
	Experience int `json:"experience"`
	Education  int `json:"education"`
}

type ATSScoreResponse struct {
	OverallScore    int          `json:"overall_score"`
	SectionScores   SectionScore `json:"section_scores"`
	MissingKeywords []string     `json:"missing_keywords"`
	Strengths       []string     `json:"strengths"`
	Weaknesses      []string     `json:"weaknesses"`
	Suggestions     []string     `json:"suggestions"`
}
