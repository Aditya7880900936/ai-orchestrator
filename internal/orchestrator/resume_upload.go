package orchestrator

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/Aditya7880900936/ai-orchestrator/internal/cache"
	"github.com/Aditya7880900936/ai-orchestrator/internal/parser/document"
)

func UploadResume(file *multipart.FileHeader) (string, error) {

	sessionID := uuid.NewString()

	src, err := file.Open()
	if err != nil {
		return "", err
	}
	defer src.Close()

	tmpFile := filepath.Join(os.TempDir(), sessionID+filepath.Ext(file.Filename))

	dst, err := os.Create(tmpFile)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", err
	}

	defer os.Remove(tmpFile)

	parser, err := document.NewParser(file.Filename)
	if err != nil {
		return "", err
	}

	text, err := parser.Parse(tmpFile)
	if err != nil {
		return "", err
	}

	if text == "" {
		return "", fmt.Errorf("resume text is empty")
	}

	if err := cache.SaveSession(sessionID, text); err != nil {
		return "", err
	}

	return sessionID, nil
}
