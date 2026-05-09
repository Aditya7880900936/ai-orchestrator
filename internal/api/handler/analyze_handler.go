package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

func AnalyzeHandler(c *gin.Context) {

	var req models.AnalyzeRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid request",
		})
		return
	}

	if req.Prompt == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "prompt required",
		})
		return
	}

	resp, err := orchestrator.Analyze(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}