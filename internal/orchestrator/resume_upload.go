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

var (
	newUUID = uuid.NewString

	openMultipartFile = func(file *multipart.FileHeader) (multipart.File, error) {
		return file.Open()
	}

	createTempFile = os.Create

	copyContent = io.Copy

	removeFile = os.Remove

	saveResumeSession = cache.SaveSession

	newDocumentParser = document.NewParser
)

func UploadResume(file *multipart.FileHeader) (string, error) {

	sessionID := newUUID()

	src, err := openMultipartFile(file)
	if err != nil {
		return "", err
	}
	defer src.Close()

	tmpFile := filepath.Join(
		os.TempDir(),
		sessionID+filepath.Ext(file.Filename),
	)

	dst, err := createTempFile(tmpFile)
	if err != nil {
		return "", err
	}
	defer dst.Close()

	if _, err := copyContent(dst, src); err != nil {
		return "", err
	}

	defer removeFile(tmpFile)

	parser, err := newDocumentParser(file.Filename)
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

	if err := saveResumeSession(sessionID, text); err != nil {
		return "", err
	}

	return sessionID, nil
}
