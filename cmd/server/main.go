package main

import (
	"log"

	handler "github.com/Aditya7880900936/ai-orchestrator/internal/api/handler"
	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	cache.InitRedis()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env")
	}

	if err := llm.InitGemini(); err != nil {
		log.Fatal(err)
	}

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
