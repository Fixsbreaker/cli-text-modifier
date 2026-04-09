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

	reQuotes := regexp.MustCompile(`'\s*(.*?)\s*'`)
	resultText = reQuotes.ReplaceAllString(resultText, "'$1'")

	words = strings.Fields(resultText) // fix indefinite articles (a to an) before vowels
	for i := 0; i < len(words)-1; i++ {
		w := words[i]
		if w == "a" || w == "A" {
			nextWord := words[i+1]
			cleanNext := strings.TrimLeft(nextWord, "'\"") // protection of
			if len(cleanNext) > 0 {
				firstChar := strings.ToLower(string([]rune(cleanNext)[0]))
				if strings.ContainsAny(firstChar, "aeiouh") {
					if w == "a" {
						words[i] = "an"
					} else {
						words[i] = "An"
					}
				}
			}
		}
	}

	return strings.Join(words, " ")
}

func applyHex(words *[]string) { // converts previous hex word to decimal
	if len(*words) == 0 {
		return
	}
	lastIdx := len(*words) - 1
	lastWord := (*words)[lastIdx]
	if val, err := strconv.ParseInt(lastWord, 16, 64); err == nil {
		(*words)[lastIdx] = strconv.FormatInt(val, 10)
	}
}

func applyBin(words *[]string) { //converts previous binary word to decimal
	if len(*words) == 0 {
		return
	}
	lastIdx := len(*words) - 1
	lastWord := (*words)[lastIdx]
	if val, err := strconv.ParseInt(lastWord, 2, 64); err == nil {
		(*words)[lastIdx] = strconv.FormatInt(val, 10)
	}
}

// applies a given text transformation to the last N words
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
