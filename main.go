package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/utils"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

func main() {

	cep := "01307-000"

	watcher := observer.Observer{
		NextHandler: func(item interface{}) {
			fmt.Printf("Processing: %v\n", item)
		},
		ErrHandler: func(err error) {
			fmt.Printf("Encountered error: %v\n", err)
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
		Subscribe(watcher)
}