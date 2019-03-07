package services

import (
	"encoding/json"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
)

// FetchViaCepService - fetch data from viacep api
func FetchViaCepService(cep string, channel chan models.Status) {
	response, fetchError := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	cepResponse := models.ViaCepResponse{}
	errorStatus := models.Status{Ok: false}

	if fetchError != nil {
		errorStatus.Value = fetchError
		channel <- errorStatus
	}

	parseHasErrors := json.NewDecoder(response.Body).Decode(&cepResponse)
	if parseHasErrors != nil {
		errorStatus.Value = parseHasErrors
		channel <- errorStatus
	}

	res := models.Status{Ok: true}
	if cepResponse.Cep == "" {
		res.Value = nil
	} else {
		res.Value = cepResponse
	}
	channel <- res

	defer response.Body.Close()
}
