package services

import (
	"net/http"

	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

// FetchViaCepService - fetch data from via cep api
func FetchViaCepService(cep string) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		response, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

		if err != nil {
			emitter.OnError(err)
		}

		emitter.OnNext(response)
		emitter.OnDone()
	})
}
