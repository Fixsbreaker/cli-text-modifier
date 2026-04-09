package processor

import (
	"regexp"
	"strconv"
	"strings"
	"unicode/utf8"
)

func Process(text string) string {
	words := strings.Fields(text)
	var processed []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		switch word {
		case "(hex)":
			applyHex(&processed)
		case "(bin)":
			applyBin(&processed)
		case "(up)":
			applyCase(&processed, strings.ToUpper, 1)
		case "(low)":
			applyCase(&processed, strings.ToLower, 1)
		case "(cap)":
			applyCase(&processed, capitalize, 1)
		case "(up,", "(low,", "(cap,":
			if i+1 < len(words) {
				numStr := words[i+1]
				if strings.HasSuffix(numStr, ")") {
					numStr = strings.TrimSuffix(numStr, ")")
					n, err := strconv.Atoi(numStr)
					if err == nil {
						switch word {
						case "(up,":
							applyCase(&processed, strings.ToUpper, n)
						case "(low,":
							applyCase(&processed, strings.ToLower, n)
						case "(cap,":
							applyCase(&processed, capitalize, n)
						}
						i++ // skip next word(num)
						continue
					}
				}
			}
			processed = append(processed, word)
		default:
			processed = append(processed, word)
		}
	}

	resultText := strings.Join(processed, " ")

	reSpaceBeforePunct := regexp.MustCompile(`\s+([.,!?:;]+)`)
	resultText = reSpaceBeforePunct.ReplaceAllString(resultText, "$1")

	reSpaceAfterPunct := regexp.MustCompile(`([.,!?:;]+)([^ .,!?:;])`)
	resultText = reSpaceAfterPunct.ReplaceAllString(resultText, "$1 $2")

	return resultText
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

func applyCase(words *[]string, transform func(string) string, n int) {
	start := len(*words) - n
	if start < 0 {
		start = 0 // protect from going out of the bounds
	}
	for i := start; i < len(*words); i++ {
		(*words)[i] = transform((*words)[i])
	}
}

func capitalize(s string) string {
	if s == "" {
		return ""
	}
	r, size := utf8.DecodeRuneInString(s)
	return strings.ToUpper(string(r)) + strings.ToLower(s[size:])
}
