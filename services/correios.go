package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
	"github.com/reactivex/rxgo/observable"
	"github.com/reactivex/rxgo/observer"
)

// FetchCepCorreiosService - fetch data from correios api
func FetchCepCorreiosService(cepRaw interface{}) observable.Observable {
	const proxyURL = "https://proxier.now.sh/"
	return observable.Create(func(emitter *observer.Observer, disposed bool) {
		cep, _ := cepRaw.(string)
		client := &http.Client{}

		url := proxyURL + "https://apps.correios.com.br/SigepMasterJPA/AtendeClienteService/AtendeCliente?wsdl"
		payload := `
			<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:cli="http://cliente.bean.master.sigep.bsb.correios.com.br/">
				<soapenv:Header/>
				<soapenv:Body>
					<cli:consultaCEP>
						<cep>` + cep + `s</cep>
					</cli:consultaCEP>
				</soapenv:Body>
			</soapenv:Envelope>
		`

		request, createRequestError := http.NewRequest("POST", url, bytes.NewBufferString(payload))
		if createRequestError != nil {
			emitter.OnError(createRequestError)
		}
		request.Header.Set("content-type", "application/soap+xml;charset=utf-8")
		request.Header.Set("cache-control", "no-cache")

		response, fetchError := client.Do(request)
		if fetchError != nil {
			emitter.OnError(fetchError)
		}
		cepResponse := models.CorreiosResponse{}
		parseHasErrors := xml.NewDecoder(response.Body).Decode(&cepResponse)
		if parseHasErrors != nil {
			emitter.OnError(errors.New("Error on parse xml"))
		}
		emitter.OnNext(cepResponse.Body.Consult.Return)
		emitter.OnDone()

		defer response.Body.Close()
	})
}
