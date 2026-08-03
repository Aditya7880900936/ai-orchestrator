package repository

import (
	"context"

	"github.com/Aditya7880900936/ai-orchestrator/internal/database"
)

type AnalysisRepository struct{}

func NewAnalysisRepository() *AnalysisRepository {
	return &AnalysisRepository{}
}

func (r *AnalysisRepository) CreateTable() error {

	query := `
	CREATE TABLE IF NOT EXISTS resume_analysis (
		id SERIAL PRIMARY KEY,
		resume TEXT NOT NULL,
		result JSONB NOT NULL,
		created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := database.DB.Exec(context.Background(), query)
	return err
}

func (r *AnalysisRepository) Save(resume string, result string) error {

	query := `
	INSERT INTO resume_analysis(resume, result)
	VALUES ($1, $2)
	`

	_, err := database.DB.Exec(
		context.Background(),
		query,
		resume,
		result,
	)

	return err
}
