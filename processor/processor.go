package processor

import (
	"strconv"
	"strings"
	"unicode/utf8"
)

func Process(text string) string {
	words := strings.Fields(text)
	var processed []string

	for _, word := range words {
		switch word {
		case "(hex)":
			applyHex(&processed)
		case "(bin)":
			applyBin(&processed)
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

func applyHex(words *[]string) {
	if len(*words) == 0 {
		return
	}
	last := len(*words) - 1
	if val, err := strconv.ParseInt((*words)[last], 16, 64); err == nil {
		(*words)[last] = strconv.FormatInt(val, 10)
	}
}

func applyBin(words *[]string) {
	if len(*words) == 0 {
		return
	}
	last := len(*words) - 1
	if val, err := strconv.ParseInt((*words)[last], 2, 64); err == nil {
		(*words)[last] = strconv.FormatInt(val, 10)
	}
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(r)) + strings.ToLower(s[size:])
}
