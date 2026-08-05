package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

// Dependency injection point
var matchJob = orchestrator.MatchJob

// MatchJob godoc
//
//	@Summary		Match Resume with Job
//	@Description	Matches a candidate's resume against a job description and returns an AI-generated job compatibility analysis.
//	@Tags			Job Matching
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.JobMatchRequest	true	"Job Match Request"
//	@Success		200		{object}	models.JobMatchResponse	"Job match analysis completed successfully"
//	@Failure		400		{object}	map[string]string		"Invalid request payload"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/job/match [post]
func MatchJob(c *gin.Context) {
	var req models.JobMatchRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := matchJob(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
