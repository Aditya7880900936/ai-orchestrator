package handler

import (
	"net/http"

	models "github.com/Aditya7880900936/ai-orchestrator/internal/model"
	"github.com/Aditya7880900936/ai-orchestrator/internal/orchestrator"
	"github.com/gin-gonic/gin"
)

// Dependency injection point
var chatWithResume = orchestrator.ChatWithResume

// ChatWithResume godoc
//
//	@Summary		Chat with Resume
//	@Description	Enables conversational AI interactions over an uploaded resume, allowing users to ask questions and receive context-aware responses.
//	@Tags			Resume
//	@Accept			json
//	@Produce		json
//	@Param			request	body		models.ResumeChatRequest	true	"Resume Chat Request"
//	@Success		200		{object}	models.ResumeChatResponse	"Resume chat response generated successfully"
//	@Failure		400		{object}	map[string]string			"Invalid request payload"
//	@Failure		500		{object}	map[string]string			"Internal server error"
//	@Router			/resume/chat [post]
func ChatWithResume(c *gin.Context) {
	var req models.ResumeChatRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	resp, err := chatWithResume(req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, resp)
}
