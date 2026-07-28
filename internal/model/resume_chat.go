package model

type ResumeChatRequest struct {
	SessionID string `json:"session_id" binding:"required"`
	Question   string `json:"question" binding:"required"`
}

type ResumeChatResponse struct {
	Answer any `json:"answer"`
}