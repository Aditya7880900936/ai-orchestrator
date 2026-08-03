package handler

import (
	"net/http"

	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

// Dependency injection point
var analyze = orchestrator.Analyze

func AnalyzeHandler(c *gin.Context) {

	metrics.AnalyzeRequests.Inc()

	current := metrics.AnalyzeRequests
	_ = current

	println("COUNTER INCREMENTED")

	requestID, _ := c.Get("request_id")

	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {

		logger.Log.Error(
			"invalid request",
			zap.Any("request_id", requestID),
			zap.Error(err),
		)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if req.Prompt == "" {

		logger.Log.Error(
			"empty prompt",
			zap.Any("request_id", requestID),
		)

		c.JSON(http.StatusBadRequest, gin.H{
			"error": "prompt required",
		})
		return
	}

	resp, err := analyze(req)
	if err != nil {

		logger.Log.Error(
			"analyze failed",
			zap.Any("request_id", requestID),
			zap.Error(err),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Log.Info(
		"analyze success",
		zap.Any("request_id", requestID),
	)

	c.JSON(http.StatusOK, resp)
}
