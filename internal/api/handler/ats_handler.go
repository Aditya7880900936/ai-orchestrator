package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

// Dependency injection point
var calculateATS = orchestrator.CalculateATS

// CalculateATS godoc
//
//	@Summary		Calculate ATS Score
//	@Description	Calculates the ATS compatibility score of a resume against the provided job description using AI analysis.
//	@Tags			ATS
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ATSScoreRequest	true	"ATS Score Request"
//	@Success		200		{object}	models.ATSScoreResponse	"ATS score calculated successfully"
//	@Failure		400		{object}	map[string]string		"Invalid request payload"
//	@Failure		500		{object}	map[string]string		"Internal server error"
//	@Router			/ats/score [post]
func CalculateATS(c *gin.Context) {
	var req models.ATSScoreRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := calculateATS(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
