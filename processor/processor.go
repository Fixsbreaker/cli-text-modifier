package processor

import (
	"strings"
	"unicode/utf8"
)

func Process(text string) string {
	words := strings.Fields(text)
	var processed []string

	for _, word := range words {
		switch word {
		case "(up)":
			if len(processed) > 0 {
				processed[len(processed)-1] = strings.ToUpper(processed[len(processed)-1])
			}
		case "(low)":
			if len(processed) > 0 {
				processed[len(processed)-1] = strings.ToLower(processed[len(processed)-1])
			}
		case "(cap)":
			if len(processed) > 0 {
				processed[len(processed)-1] = capitalize(processed[len(processed)-1])
			}
		default:
			processed = append(processed, word)
		}
	}

	return strings.Join(processed, " ")
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(r)) + strings.ToLower(s[size:])
}
