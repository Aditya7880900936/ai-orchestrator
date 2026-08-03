package handler

import (
	"bytes"
	"errors"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestUploadResume_MissingFile(t *testing.T) {

	setupHandlerTest(t)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	body := &bytes.Buffer{}

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/upload",
		body,
	)

	req.Header.Set(
		"Content-Type",
		"multipart/form-data",
	)

	c.Request = req

	UploadResume(c)

	if w.Code != http.StatusBadRequest {
		t.Fatalf("expected %d got %d", http.StatusBadRequest, w.Code)
	}

	if !strings.Contains(w.Body.String(), "resume is required") {
		t.Fatal("expected error message")
	}
}

func TestUploadResume_Error(t *testing.T) {

	setupHandlerTest(t)

	old := uploadResume
	defer func() {
		uploadResume = old
	}()

	uploadResume = func(file *multipart.FileHeader) (string, error) {

		if file == nil {
			t.Fatal("expected file")
		}

		return "", errors.New("upload failed")
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(
		"resume",
		"resume.pdf",
	)
	if err != nil {
		t.Fatal(err)
	}

	_, _ = io.WriteString(part, "dummy resume")

	writer.Close()

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/upload",
		body,
	)

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	UploadResume(c)

	if w.Code != http.StatusInternalServerError {
		t.Fatalf(
			"expected %d got %d",
			http.StatusInternalServerError,
			w.Code,
		)
	}

	if !strings.Contains(w.Body.String(), "upload failed") {
		t.Fatal("expected upload failed")
	}
}

func TestUploadResume_Success(t *testing.T) {

	setupHandlerTest(t)

	old := uploadResume
	defer func() {
		uploadResume = old
	}()

	uploadResume = func(file *multipart.FileHeader) (string, error) {

		if file == nil {
			t.Fatal("expected file")
		}

		if file.Filename != "resume.pdf" {
			t.Fatalf(
				"unexpected filename %s",
				file.Filename,
			)
		}

		return "session-123", nil
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)

	part, err := writer.CreateFormFile(
		"resume",
		"resume.pdf",
	)
	if err != nil {
		t.Fatal(err)
	}

	_, _ = io.WriteString(part, "dummy resume")

	writer.Close()

	req, _ := http.NewRequest(
		http.MethodPost,
		"/resume/upload",
		body,
	)

	req.Header.Set(
		"Content-Type",
		writer.FormDataContentType(),
	)

	w := httptest.NewRecorder()
	c, _ := gin.CreateTestContext(w)

	c.Request = req

	UploadResume(c)

	if w.Code != http.StatusOK {
		t.Fatalf(
			"expected %d got %d",
			http.StatusOK,
			w.Code,
		)
	}

	resp := w.Body.String()

	if !strings.Contains(resp, "session-123") {
		t.Fatal("expected session id")
	}

	if !strings.Contains(resp, "Resume uploaded successfully") {
		t.Fatal("expected success message")
	}
}