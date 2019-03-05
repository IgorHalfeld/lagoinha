package services

import (
	"encoding/json"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

const proxyURL = "https://proxier.now.sh/"
const token = "37d718d2984e6452584a76d3d59d3a26"

// FetchCepAbertoService - fetch data from cepaberto api
func FetchCepAbertoService(cepRaw interface{}) observable.Observable {
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		client := &http.Client{}

		url := proxyURL + "http://www.cepaberto.com/api/v2/ceps.json?cep=" + cep
		request, createRequestError := http.NewRequest("GET", url, nil)
		if createRequestError != nil {
			emitter.OnError(createRequestError)
		}
		request.Header.Set("content-type", "application/json;charset=utf-8")
		request.Header.Set("Authorization", "Token token="+token)

		response, fetchError := client.Do(request)
		if fetchError != nil {
			emitter.OnError(fetchError)
		}
		cepResponse := models.CepAbertoResponse{}
		parseHasErrors := json.NewDecoder(response.Body).Decode(&cepResponse)
		if parseHasErrors != nil {
			emitter.OnError(parseHasErrors)
		}
		emitter.OnNext(cepResponse)
		emitter.OnDone()

		defer response.Body.Close()
	})
}
