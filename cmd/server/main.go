package main

import (
	"log"
	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	llm.InitOpenAI()

	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": "ok",
		})
	})

	r.POST("/analyze", handler.AnalyzeHandler)

	log.Println("Server running on :8080")
	r.Run(":8080")
}