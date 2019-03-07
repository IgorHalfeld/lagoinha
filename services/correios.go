package services

import (
	"bytes"
	"encoding/xml"
	"net/http"

	"github.com/igorhalfeld/lagoinha/models"
)

// FetchCepCorreiosService - fetch data from correios api
func FetchCepCorreiosService(cep string, channel chan models.Status) {
	const proxyURL = "https://proxier.now.sh/"
	client := &http.Client{}
	cepResponse := models.CorreiosResponse{}
	errorStatus := models.Status{Ok: false}

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
		errorStatus.Value = createRequestError
		channel <- errorStatus
	}
	request.Header.Set("content-type", "application/soap+xml;charset=utf-8")
	request.Header.Set("cache-control", "no-cache")

	response, fetchError := client.Do(request)
	if fetchError != nil {
		errorStatus.Value = fetchError
		channel <- errorStatus
	}

	parseHasErrors := xml.NewDecoder(response.Body).Decode(&cepResponse)
	if parseHasErrors != nil {
		errorStatus.Value = parseHasErrors
		channel <- errorStatus
	}

	res := models.Status{Ok: true}
	if cepResponse.Body.Consult.Return.Cep == "" {
		res.Value = nil
	} else {
		res.Value = cepResponse.Body.Consult.Return
	}
	channel <- res

	defer response.Body.Close()
}
