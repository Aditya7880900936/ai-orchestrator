package model

type ResumeChatRequest struct {
	SessionID string `json:"session_id" binding:"required"`
	ResumeText string `json:"resume_text,omitempty"`
	Question   string `json:"question" binding:"required"`
}

type ResumeChatResponse struct {
	Answer string `json:"answer"`
}