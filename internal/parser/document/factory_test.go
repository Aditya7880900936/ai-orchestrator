package document

import "testing"

func TestNewParser_PDF(t *testing.T) {

	p, err := NewParser("resume.pdf")

	if err != nil {
		t.Fatal(err)
	}

	if _, ok := p.(*PDFParser); !ok {
		t.Fatal("expected PDFParser")
	}
}

func TestNewParser_DOCX(t *testing.T) {

	p, err := NewParser("resume.docx")

	if err != nil {
		t.Fatal(err)
	}

	if _, ok := p.(*DOCXParser); !ok {
		t.Fatal("expected DOCXParser")
	}
}

func TestNewParser_CaseInsensitive(t *testing.T) {

	p, err := NewParser("resume.PDF")

	if err != nil {
		t.Fatal(err)
	}

	if _, ok := p.(*PDFParser); !ok {
		t.Fatal("expected PDFParser")
	}
}

func TestNewParser_Unsupported(t *testing.T) {

	p, err := NewParser("resume.txt")

	if err == nil {
		t.Fatal("expected error")
	}

	if p != nil {
		t.Fatal("expected nil parser")
	}
}
