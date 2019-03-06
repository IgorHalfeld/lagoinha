package services

import (
	"encoding/json"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
)

// FetchCepAbertoService - fetch data from cepaberto api
func FetchCepAbertoService(cep string, channel chan models.Status) {
	const proxyURL = "https://proxier.now.sh/"
	const token = "37d718d2984e6452584a76d3d59d3a26"
	client := &http.Client{}
	cepResponse := models.CepAbertoResponse{}
	errorStatus := models.Status{Ok: false}

	url := proxyURL + "http://www.cepaberto.com/api/v2/ceps.json?cep=" + cep
	request, createRequestError := http.NewRequest("GET", url, nil)
	if createRequestError != nil {
		errorStatus.Value = createRequestError
		channel <- errorStatus
	}
	request.Header.Set("content-type", "application/json;charset=utf-8")
	request.Header.Set("Authorization", "Token token="+token)

	response, fetchError := client.Do(request)
	if fetchError != nil {
		errorStatus.Value = fetchError
		channel <- errorStatus
	}

	parseHasErrors := json.NewDecoder(response.Body).Decode(&cepResponse)
	if parseHasErrors != nil {
		errorStatus.Value = parseHasErrors
		channel <- errorStatus
	}
	channel <- models.Status{
		Ok:    true,
		Value: cepResponse,
	}

	defer response.Body.Close()
}
