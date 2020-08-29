package utils

import (
	"regexp"
	"strings"
)

// ValidateInputLength Validate input length
func ValidateInputLength(cepRaw string) (status bool) {
	const cepSize = 8
	cepLength := len(cepRaw)
	if cepLength <= cepSize {
		status = true
	} else {
		status = false
	}
	return status
}

// RemoveSpecialCharacters Remove special characters
func RemoveSpecialCharacters(cepRaw string) (cepParsed string) {
	rule := regexp.MustCompile(`\D+`)
	cepParsed = rule.ReplaceAllString(cepRaw, "")
	return cepParsed
}

// LeftPadWithZeros Pad cep with zeros
func LeftPadWithZeros(cepRaw string) (cepParsed string) {
	const cepSize = 8
	cepLength := len(cepRaw)
	timesToRepeat := cepSize - cepLength
	pad := strings.Repeat("0", timesToRepeat)
	cepParsed = pad + cepRaw
	return cepParsed
}
