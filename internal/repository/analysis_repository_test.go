package repository

import (
	"testing"

	"github.com/Aditya7880900936/ai-orchestrator/internal/database"
)

func TestCreateTable(t *testing.T) {

	database.Init()

	repo := NewAnalysisRepository()

	if err := repo.CreateTable(); err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}
}

func TestSave(t *testing.T) {

	database.Init()

	repo := NewAnalysisRepository()

	if err := repo.CreateTable(); err != nil {
		t.Fatalf("CreateTable failed: %v", err)
	}

	err := repo.Save(
		"My Resume",
		`{"summary":"Backend Engineer"}`,
	)

	if err != nil {
		t.Fatalf("Save failed: %v", err)
	}
}
