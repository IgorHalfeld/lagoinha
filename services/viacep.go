package services

import (
	"encoding/json"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

// FetchViaCepService - fetch data from viacep api
func FetchViaCepService(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		response, fetchError := http.Get("https://viacep.com.br/ws/" + cep + "/json/")

		if fetchError != nil {
			emitter.OnError(fetchError)
		}

		cepResponse := models.ViaCepResponse{}
		parseHasErrors := json.NewDecoder(response.Body).Decode(&cepResponse)
		if parseHasErrors != nil {
			emitter.OnError(parseHasErrors)
		}
		emitter.OnNext(cepResponse)

		defer response.Body.Close()
		emitter.OnDone()
	})
}
