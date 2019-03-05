package utils

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

const cepSize = 8

// ValidateInputType - Validate input cep
func ValidateInputType(cepRaw interface{}) interface{} {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {

		_, isString := cepRaw.(string)

		if isString {
			emitter.OnNext(cepRaw)
		} else {
			emitter.OnError(errors.New("Cep must be string"))
		}
	})
}

// ValidateInputLength - Validate input length
func ValidateInputLength(cepRaw interface{}) interface{} {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		cepLength := utf8.RuneCountInString(cep)
		if cepLength <= cepSize {
			emitter.OnNext(cep)
		} else {
			emitter.OnError(errors.New("Cep length is less than 8 characters"))
		}
	})
}

// RemoveSpecialCharacters - Remove special characters
func RemoveSpecialCharacters(cepRaw string) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		rule := regexp.MustCompile(`/\D+/g`)
		cleanCep := rule.ReplaceAllString(cepRaw, "")
		emitter.OnNext(cleanCep)
	})
}

// LeftPadWithZeros - Pad cep with zeros
func LeftPadWithZeros(cepRaw string) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cepLength := utf8.RuneCountInString(cepRaw)
		timesToRepeat := cepSize - cepLength
		pad := strings.Repeat("0", timesToRepeat)
		emitter.OnNext(pad + cepRaw)
	})
}
