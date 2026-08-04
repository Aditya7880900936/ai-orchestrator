package document

import "testing"

func TestDOCXParser_InvalidFile(t *testing.T) {

	p := &DOCXParser{}

	_, err := p.Parse("does-not-exist.docx")

	if err == nil {
		t.Fatal("expected error")
	}
}
