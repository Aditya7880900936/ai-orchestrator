package model

type ResumeUploadResponse struct {
	SessionID string `json:"session_id"`
	Message   string `json:"message"`
}
