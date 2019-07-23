package utils

import (
	"fmt"
	"regexp"
	"strings"

	"github.com/igorhalfeld/lagoinha/models"
	"github.com/igorhalfeld/lagoinha/services"
)

// ValidateInputLength - Validate input length
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

// RemoveSpecialCharacters - Remove special characters
func RemoveSpecialCharacters(cepRaw string) (cepParsed string) {
	rule := regexp.MustCompile(`\D+`)
	cepParsed = rule.ReplaceAllString(cepRaw, "")
	return cepParsed
}

// LeftPadWithZeros - Pad cep with zeros
func LeftPadWithZeros(cepRaw string) (cepParsed string) {
	const cepSize = 8
	cepLength := len(cepRaw)
	timesToRepeat := cepSize - cepLength
	pad := strings.Repeat("0", timesToRepeat)
	cepParsed = pad + cepRaw
	return cepParsed
}

// RaceServices - run parallel and return first result
func RaceServices(cepRaw string) (value interface{}, err error) {
	response := make(chan models.Status)
	go services.FetchCepCorreiosService(cepRaw, response)
	go services.FetchViaCepService(cepRaw, response)
	go services.FetchCepAbertoService(cepRaw, response)
	for {
		status := <-response
		if status.Ok {
			value = status.Value
			return value, nil
		}

		return nil, fmt.Errorf("Error on fetch address from cep %s", status.Value)
	}
}
