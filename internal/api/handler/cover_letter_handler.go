package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

// Dependency injection point
var generateCoverLetter = orchestrator.GenerateCoverLetter

// GenerateCoverLetter godoc
//
//	@Summary		Generate Cover Letter
//	@Description	Generates a personalized AI-powered cover letter based on the provided resume and job description.
//	@Tags			Cover Letter
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.CoverLetterRequest	true	"Cover Letter Request"
//	@Success		200		{object}	models.CoverLetterResponse	"Cover letter generated successfully"
//	@Failure		400		{object}	map[string]string			"Invalid request payload"
//	@Failure		500		{object}	map[string]string			"Internal server error"
//	@Router			/cover-letter/generate [post]
func GenerateCoverLetter(c *gin.Context) {
	var req models.CoverLetterRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := generateCoverLetter(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
