package model

type ResumeImproveRequest struct {
	ResumeText string `json:"resume_text" binding:"required"`
}

type ResumeImproveResponse struct {
	ImprovedSummary     string   `json:"improved_summary"`
	ImprovedExperience  []string `json:"improved_experience"`
	ImprovedProjects    []string `json:"improved_projects"`
	MissingSections     []string `json:"missing_sections"`
	ActionVerbs         []string `json:"action_verbs"`
	OverallSuggestions  []string `json:"overall_suggestions"`
}