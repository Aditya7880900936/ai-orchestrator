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

func AnalyzeResume(c *gin.Context) {

	var req model.ResumeAnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		logger.Log.Error("invalid resume request", zap.Error(err))
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	requestID, _ := c.Get("request_id")

	metrics.AnalyzeRequests.Inc()

	res, err := orchestrator.AnalyzeResume(req)
	if err != nil {
		logger.Log.Error("resume analysis failed",
			zap.Error(err),
			zap.Any("request_id", requestID),
		)

		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	logger.Log.Info("resume analysis success",
		zap.Any("request_id", requestID),
	)

	c.JSON(http.StatusOK, res)
}