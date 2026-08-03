package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"

	"github.com/gin-gonic/gin"
)

// Dependency injection point
var uploadResume = orchestrator.UploadResume

func UploadResume(c *gin.Context) {

	file, err := c.FormFile("resume")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "resume is required",
		})
		return
	}

	sessionID, err := uploadResume(file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, models.ResumeUploadResponse{
		SessionID: sessionID,
		Message:   "Resume uploaded successfully",
	})
}
