package document

import (
	"fmt"
	"path/filepath"
	"strings"
)

func NewParser(filename string) (Parser, error) {
	ext := strings.ToLower(filepath.Ext(filename))

	switch ext {
	case ".pdf":
		return &PDFParser{}, nil

	case ".docx":
		return &DOCXParser{}, nil

	default:
		return nil, fmt.Errorf("unsupported file type: %s", ext)
	}
}
