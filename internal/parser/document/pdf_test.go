package document

import "testing"

func TestPDFParser_InvalidFile(t *testing.T) {

	p := &PDFParser{}

	_, err := p.Parse("does-not-exist.pdf")

	if err == nil {
		t.Fatal("expected error")
	}
}
