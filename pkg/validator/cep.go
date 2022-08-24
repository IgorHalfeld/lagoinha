package validator

import (
	"regexp"
	"strings"
)

// ValidateInputLength validate input length
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

// RemoveSpecialCharacters remove special characters
func RemoveSpecialCharacters(cepRaw string) (cepParsed string) {
	rule := regexp.MustCompile(`\D+`)
	cepParsed = rule.ReplaceAllString(cepRaw, "")
	return cepParsed
}

// LeftPadWithZeros pad cep with zeros
func LeftPadWithZeros(cepRaw string) (cepParsed string) {
	const cepSize = 8
	cepLength := len(cepRaw)
	timesToRepeat := cepSize - cepLength
	pad := strings.Repeat("0", timesToRepeat)
	cepParsed = pad + cepRaw
	return cepParsed
}
