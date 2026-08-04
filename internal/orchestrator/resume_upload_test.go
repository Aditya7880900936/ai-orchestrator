package orchestrator

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"os"
	"testing"

	"github.com/Aditya7880900936/ai-orchestrator/internal/parser/document"
)

type mockParser struct {
	parse func(path string) (string, error)
}

func (m *mockParser) Parse(path string) (string, error) {
	return m.parse(path)
}

func createMultipartFile(t *testing.T) *multipart.FileHeader {

	var body bytes.Buffer

	writer := multipart.NewWriter(&body)

	part, err := writer.CreateFormFile("resume", "resume.pdf")
	if err != nil {
		t.Fatal(err)
	}

	_, err = part.Write([]byte("dummy pdf"))
	if err != nil {
		t.Fatal(err)
	}

	writer.Close()

	req := multipart.NewReader(
		&body,
		writer.Boundary(),
	)

	form, err := req.ReadForm(1024 * 1024)
	if err != nil {
		t.Fatal(err)
	}

	return form.File["resume"][0]
}

func TestUploadResume_Success(t *testing.T) {

	oldUUID := newUUID
	oldParser := newDocumentParser
	oldSave := saveResumeSession

	defer func() {
		newUUID = oldUUID
		newDocumentParser = oldParser
		saveResumeSession = oldSave
	}()

	newUUID = func() string {
		return "session-123"
	}

	newDocumentParser = func(name string) (document.Parser, error) {
		return &mockParser{
			parse: func(path string) (string, error) {
				return "Backend Resume", nil
			},
		}, nil
	}

	saved := false

	saveResumeSession = func(id, text string) error {

		saved = true

		if id != "session-123" {
			t.Fatal("wrong session id")
		}

		if text != "Backend Resume" {
			t.Fatal("wrong resume")
		}

		return nil
	}

	file := createMultipartFile(t)

	sessionID, err := UploadResume(file)
	if err != nil {
		t.Fatal(err)
	}

	if sessionID != "session-123" {
		t.Fatal("wrong session id")
	}

	if !saved {
		t.Fatal("session not saved")
	}
}

func TestUploadResume_ParserError(t *testing.T) {

	oldParser := newDocumentParser

	defer func() {
		newDocumentParser = oldParser
	}()

	newDocumentParser = func(name string) (document.Parser, error) {
		return nil, errors.New("parser failed")
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUploadResume_ParseError(t *testing.T) {

	oldParser := newDocumentParser

	defer func() {
		newDocumentParser = oldParser
	}()

	newDocumentParser = func(name string) (document.Parser, error) {
		return &mockParser{
			parse: func(path string) (string, error) {
				return "", errors.New("parse failed")
			},
		}, nil
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUploadResume_EmptyResume(t *testing.T) {

	oldParser := newDocumentParser

	defer func() {
		newDocumentParser = oldParser
	}()

	newDocumentParser = func(name string) (document.Parser, error) {
		return &mockParser{
			parse: func(path string) (string, error) {
				return "", nil
			},
		}, nil
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}

	if err.Error() != "resume text is empty" {
		t.Fatal(err)
	}
}

func TestUploadResume_SaveSessionError(t *testing.T) {

	oldParser := newDocumentParser
	oldSave := saveResumeSession

	defer func() {
		newDocumentParser = oldParser
		saveResumeSession = oldSave
	}()

	newDocumentParser = func(name string) (document.Parser, error) {
		return &mockParser{
			parse: func(path string) (string, error) {
				return "resume", nil
			},
		}, nil
	}

	saveResumeSession = func(id, text string) error {
		return errors.New("redis failed")
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUploadResume_OpenError(t *testing.T) {

	oldOpen := openMultipartFile

	defer func() {
		openMultipartFile = oldOpen
	}()

	openMultipartFile = func(file *multipart.FileHeader) (multipart.File, error) {
		return nil, errors.New("open failed")
	}

	file := &multipart.FileHeader{}

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUploadResume_CreateFileError(t *testing.T) {

	oldCreate := createTempFile

	defer func() {
		createTempFile = oldCreate
	}()

	createTempFile = func(name string) (*os.File, error) {
		return nil, errors.New("create failed")
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}

func TestUploadResume_CopyError(t *testing.T) {

	oldCopy := copyContent

	defer func() {
		copyContent = oldCopy
	}()

	copyContent = func(dst io.Writer, src io.Reader) (int64, error) {
		return 0, errors.New("copy failed")
	}

	file := createMultipartFile(t)

	_, err := UploadResume(file)

	if err == nil {
		t.Fatal("expected error")
	}
}
