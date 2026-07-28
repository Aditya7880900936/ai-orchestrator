package model

type CoverLetterRequest struct {
	Name           string `json:"name" binding:"required"`
	Company        string `json:"company" binding:"required"`
	Position       string `json:"position" binding:"required"`
	ResumeText     string `json:"resume_text" binding:"required"`
	JobDescription string `json:"job_description" binding:"required"`
}

type CoverLetterResponse struct {
	CoverLetter string `json:"cover_letter"`
}
