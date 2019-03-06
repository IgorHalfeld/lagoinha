package utils

import (
	"errors"
	"fmt"

	"github.com/igorhalfeld/lagoinha/models"

	"github.com/igorhalfeld/lagoinha/services"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

// RaceServices - run parallel and return first result
func RaceServices(cepRaw interface{}) observable.Observable {
	response := make(chan models.Status)

	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		go services.FetchCepCorreiosService(cep, response)
		go services.FetchViaCepService(cep, response)
		go services.FetchCepAbertoService(cep, response)

		for {
			status := <-response
			if status.Ok {
				fmt.Printf("Responsive is %v\n", status.Value)
				emitter.OnNext(status.Value)
				emitter.OnDone()
			} else {
				emitter.OnError(errors.New("Something happen"))
			}
		}
	})
}
