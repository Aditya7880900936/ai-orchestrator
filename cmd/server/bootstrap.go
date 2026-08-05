package main

import (
	"log"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	"github.com/Aditya7880900936/ai-orchestrator/internal/database"
	llm "github.com/Aditya7880900936/ai-orchestrator/internal/llm"
	"github.com/Aditya7880900936/ai-orchestrator/internal/logger"
	"github.com/Aditya7880900936/ai-orchestrator/internal/metrics"
	"github.com/Aditya7880900936/ai-orchestrator/internal/repository"

	"github.com/joho/godotenv"
)

type analysisRepository interface {
	CreateTable() error
}

// Dependency injection points
var (
	loadEnv = func() error {
		return godotenv.Load()
	}

	initRedis = cache.InitRedis

	initDatabase = database.Init

	newAnalysisRepository = func() analysisRepository {
		return repository.NewAnalysisRepository()
	}

	initMetrics = metrics.Init

	initLogger = logger.Init

	initGemini = llm.InitGemini
)

func initializeImpl() error {

	if err := loadEnv(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	initRedis()

	initDatabase()

	repo := newAnalysisRepository()

	if err := repo.CreateTable(); err != nil {
		return err
	}

	initMetrics()

	if err := initLogger(); err != nil {
		return err
	}

	if err := initGemini(); err != nil {
		return err
	}

	return nil
}

var initialize = initializeImpl
