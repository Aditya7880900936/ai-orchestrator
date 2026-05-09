package models

type AnalyzeRequest struct {
	Prompt string `json:"prompt"`
}

type AnalyzeResponse struct {
	Summary  string   `json:"summary"`
	Keywords []string `json:"keywords"`
}