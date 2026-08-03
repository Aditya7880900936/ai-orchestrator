package parser

import "strings"

func ExtractJSON(raw string) string {

	start := strings.Index(raw, "{")
	end := strings.LastIndex(raw, "}")

	if start == -1 || end == -1 || start >= end {
		return ""
	}

	return raw[start : end+1]
}
