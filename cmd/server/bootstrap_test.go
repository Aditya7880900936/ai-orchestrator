package main

import (
	"errors"
	"testing"
)

type mockRepository struct {
	err error
}

func (m *mockRepository) CreateTable() error {
	return m.err
}

func TestInitialize_Success(t *testing.T) {

	oldLoad := loadEnv
	oldRedis := initRedis
	oldDB := initDatabase
	oldRepo := newAnalysisRepository
	oldMetrics := initMetrics
	oldLogger := initLogger
	oldGemini := initGemini

	defer func() {
		loadEnv = oldLoad
		initRedis = oldRedis
		initDatabase = oldDB
		newAnalysisRepository = oldRepo
		initMetrics = oldMetrics
		initLogger = oldLogger
		initGemini = oldGemini
	}()

	calls := 0

	loadEnv = func() error {
		calls++
		return nil
	}

	initRedis = func() {
		calls++
	}

	initDatabase = func() {
		calls++
	}

	newAnalysisRepository = func() analysisRepository {
		calls++
		return &mockRepository{}
	}

	initMetrics = func() {
		calls++
	}

	initLogger = func() error {
		calls++
		return nil
	}

	initGemini = func() error {
		calls++
		return nil
	}

	err := initialize()

	if err != nil {
		t.Fatal(err)
	}

	if calls != 7 {
		t.Fatalf("expected 7 calls got %d", calls)
	}
}

func TestInitialize_LoadEnvError(t *testing.T) {

	oldLoad := loadEnv
	oldRedis := initRedis
	oldDB := initDatabase
	oldRepo := newAnalysisRepository
	oldMetrics := initMetrics
	oldLogger := initLogger
	oldGemini := initGemini

	defer func() {
		loadEnv = oldLoad
		initRedis = oldRedis
		initDatabase = oldDB
		newAnalysisRepository = oldRepo
		initMetrics = oldMetrics
		initLogger = oldLogger
		initGemini = oldGemini
	}()

	// Simulate missing .env
	loadEnv = func() error {
		return errors.New("env failed")
	}

	initRedis = func() {}
	initDatabase = func() {}

	newAnalysisRepository = func() analysisRepository {
		return &mockRepository{}
	}

	initMetrics = func() {}

	initLogger = func() error {
		return nil
	}

	initGemini = func() error {
		return nil
	}

	err := initialize()

	// Missing .env should NOT fail initialization
	if err != nil {
		t.Fatalf("expected initialize to continue without .env, got: %v", err)
	}
}

func TestInitialize_CreateTableError(t *testing.T) {

	oldLoad := loadEnv
	oldRedis := initRedis
	oldDB := initDatabase
	oldRepo := newAnalysisRepository
	oldMetrics := initMetrics

	defer func() {
		loadEnv = oldLoad
		initRedis = oldRedis
		initDatabase = oldDB
		newAnalysisRepository = oldRepo
		initMetrics = oldMetrics
	}()

	loadEnv = func() error { return nil }

	initRedis = func() {}

	initDatabase = func() {}

	initMetrics = func() {}

	newAnalysisRepository = func() analysisRepository {
		return &mockRepository{
			err: errors.New("create failed"),
		}
	}

	err := initialize()

	if err == nil {
		t.Fatal("expected create table error")
	}
}

func TestInitialize_LoggerError(t *testing.T) {

	oldLoad := loadEnv
	oldRedis := initRedis
	oldDB := initDatabase
	oldRepo := newAnalysisRepository
	oldMetrics := initMetrics
	oldLogger := initLogger

	defer func() {
		loadEnv = oldLoad
		initRedis = oldRedis
		initDatabase = oldDB
		newAnalysisRepository = oldRepo
		initMetrics = oldMetrics
		initLogger = oldLogger
	}()

	loadEnv = func() error { return nil }

	initRedis = func() {}

	initDatabase = func() {}

	initMetrics = func() {}

	newAnalysisRepository = func() analysisRepository {
		return &mockRepository{}
	}

	initLogger = func() error {
		return errors.New("logger failed")
	}

	err := initialize()

	if err == nil {
		t.Fatal("expected logger error")
	}
}

func TestInitialize_GeminiError(t *testing.T) {

	oldLoad := loadEnv
	oldRedis := initRedis
	oldDB := initDatabase
	oldRepo := newAnalysisRepository
	oldMetrics := initMetrics
	oldLogger := initLogger
	oldGemini := initGemini

	defer func() {
		loadEnv = oldLoad
		initRedis = oldRedis
		initDatabase = oldDB
		newAnalysisRepository = oldRepo
		initMetrics = oldMetrics
		initLogger = oldLogger
		initGemini = oldGemini
	}()

	loadEnv = func() error { return nil }

	initRedis = func() {}

	initDatabase = func() {}

	initMetrics = func() {}

	newAnalysisRepository = func() analysisRepository {
		return &mockRepository{}
	}

	initLogger = func() error {
		return nil
	}

	initGemini = func() error {
		return errors.New("gemini failed")
	}

	err := initialize()

	if err == nil {
		t.Fatal("expected gemini error")
	}
}
