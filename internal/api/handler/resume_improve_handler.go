package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

// Dependency injection point
var improveResume = orchestrator.ImproveResume

// ImproveResume godoc
//
//	@Summary		Improve Resume
//	@Description	Generates AI-powered suggestions to improve a resume.
//	@Tags			Resume
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ResumeImproveRequest	true	"Resume Improvement Request"
//	@Success		200		{object}	models.ResumeImproveResponse
//	@Failure		400		{object}	map[string]string
//	@Failure		500		{object}	map[string]string
//	@Router			/resume/improve [post]
func ImproveResume(c *gin.Context) {
	var req models.ResumeImproveRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := improveResume(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
