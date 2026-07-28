package document

import (
	"strings"

	"github.com/nguyenthenguyen/docx"
)

type DOCXParser struct{}

func (p *DOCXParser) Parse(path string) (string, error) {

	d, err := docx.ReadDocxFile(path)
	if err != nil {
		return "", err
	}
	defer d.Close()

	doc := d.Editable()

	text := strings.TrimSpace(doc.GetContent())

	return text, nil
}