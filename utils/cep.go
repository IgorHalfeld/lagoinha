package utils

import (
	"errors"
	"regexp"
	"strings"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

const cepSize = 8

// ValidateInputType - Validate input cep
func ValidateInputType(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {

		_, isString := cepRaw.(string)

		if isString {
			emitter.OnNext(cepRaw)
			emitter.OnDone()
		} else {
			emitter.OnError(errors.New("Cep must be string"))
		}
	})
}

// ValidateInputLength - Validate input length
func ValidateInputLength(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		cepLength := len(cep)
		if cepLength <= cepSize {
			emitter.OnNext(cep)
			emitter.OnDone()
		} else {
			emitter.OnError(errors.New("Cep length is less than 8 characters"))
		}
	})
}

// RemoveSpecialCharacters - Remove special characters
func RemoveSpecialCharacters(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		rule := regexp.MustCompile(`\D+`)
		cleanCep := rule.ReplaceAllString(cep, "")
		emitter.OnNext(cleanCep)
		emitter.OnDone()
	})
}

// LeftPadWithZeros - Pad cep with zeros
func LeftPadWithZeros(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		cepLength := len(cep)
		timesToRepeat := cepSize - cepLength
		pad := strings.Repeat("0", timesToRepeat)
		emitter.OnNext(pad + cep)
		emitter.OnDone()
	})
}
