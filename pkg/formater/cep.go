package formater

import (
	"regexp"
	"strings"
)

func RemoveSpecialCharacters(cepRaw string) (cepParsed string) {
	rule := regexp.MustCompile(`\D+`)
	cepParsed = rule.ReplaceAllString(cepRaw, "")
	return cepParsed
}

func LeftPadWithZeros(cepRaw string) (cepParsed string) {
	const cepSize = 8
	cepLength := len(cepRaw)
	timesToRepeat := cepSize - cepLength
	pad := strings.Repeat("0", timesToRepeat)
	cepParsed = pad + cepRaw
	return cepParsed
}
