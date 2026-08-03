package parser

import "testing"

func TestExtractJSON_Valid(t *testing.T) {

	raw := `{"name":"Aditya"}`

	got := ExtractJSON(raw)

	if got != raw {
		t.Fatalf("expected %q got %q", raw, got)
	}
}

func TestExtractJSON_WithExtraText(t *testing.T) {

	raw := `Hello user

{"summary":"Backend"}

Thank you`

	got := ExtractJSON(raw)

	expected := `{"summary":"Backend"}`

	if got != expected {
		t.Fatalf("expected %q got %q", expected, got)
	}
}

func TestExtractJSON_InvalidBraces(t *testing.T) {

	raw := `}hello{`

	got := ExtractJSON(raw)

	if got != "" {
		t.Fatalf("expected empty string got %q", got)
	}
}
