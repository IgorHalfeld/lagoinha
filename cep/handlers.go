package cep

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/utils"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

// Cep - get address
func Cep(cep string) (interface{}, interface{}) {
	var address interface{}
	var err interface{}

	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			address = item
		},
		ErrHandler: func(erro error) {
			err = erro
		},
		DoneHandler: func() {
			fmt.Println("Done!")
		},
	}

	<-observable.
		Just(cep).
		FlatMap(utils.ValidateInputType, 1).
		FlatMap(utils.RemoveSpecialCharacters, 1).
		FlatMap(utils.ValidateInputLength, 1).
		FlatMap(utils.LeftPadWithZeros, 1).
		FlatMap(utils.RaceServices, 1).
		Subscribe(watcher)

	return address, err
}
