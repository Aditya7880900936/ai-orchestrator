package handler

import (
	"net/http"

	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	"github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Dependency injection point
var extractSkills = orchestrator.ExtractSkills

// ExtractSkills godoc
//
//	@Summary		Extract Skills
//	@Description	Extracts technical and soft skills from a resume using AI-powered analysis.
//	@Tags			Skills
//	@Accept			json
//	@Produce		json
//	@Param			request	body		model.SkillExtractionRequest	true	"Skill Extraction Request"
//	@Success		200		{object}	model.SkillExtractionResponse	"Skills extracted successfully"
//	@Failure		400		{object}	map[string]string			"Invalid request payload"
//	@Failure		500		{object}	map[string]string			"Internal server error"
//	@Router			/skills/extract [post]
func ExtractSkills(c *gin.Context) {

	var req model.SkillExtractionRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("invalid skill extraction request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	requestID, _ := c.Get("request_id")

	metrics.AnalyzeRequests.Inc()

	res, err := extractSkills(req)
	if err != nil {
		logger.Log.Error("skill extraction failed",
			zap.Error(err),
			zap.Any("request_id", requestID),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Log.Info("skill extraction success",
		zap.Any("request_id", requestID),
	)

	c.JSON(http.StatusOK, res)
}
