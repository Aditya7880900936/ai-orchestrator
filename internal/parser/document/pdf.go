package document

import (
	"bytes"
	"io"

	"github.com/ledongthuc/pdf"
)

type PDFParser struct{}

func (p *PDFParser) Parse(path string) (string, error) {

	f, reader, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer

	r, err := reader.GetPlainText()
	if err != nil {
		return "", err
	}

	_, err = io.Copy(&buf, r)
	if err != nil {
		return "", err
	}

	return buf.String(), nil
}
