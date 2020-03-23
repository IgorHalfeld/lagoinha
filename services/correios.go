package services

import (
	"bytes"
	"encoding/xml"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/structs"
)

// CorreiosService service
type CorreiosService struct{}

// NewCorreiosService creates a new instance
func NewCorreiosService() *CorreiosService {
	return &CorreiosService{}
}

// Request - fetch data from correios api
func (cs *CorreiosService) Request(cep string) (*structs.Cep, error) {
	const proxyURL = "https://proxier.now.sh/"
	client := &http.Client{}

	result := structs.CorreiosResponse{}

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
	req, err := http.NewRequest("POST", url, bytes.NewBufferString(payload))
	if err != nil {
		return nil, err
	}

	req.Header.Set("content-type", "application/soap+xml;charset=utf-8")
	req.Header.Set("cache-control", "no-cache")

	res, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = xml.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return cs.formater(&result)
}

func (cs *CorreiosService) formater(r *structs.CorreiosResponse) (*structs.Cep, error) {
	if r == nil {
		return nil, errors.New("Cep not found")
	}

	cep := &structs.Cep{
		Cep:          r.Body.Consult.Return.Cep,
		City:         r.Body.Consult.Return.City,
		Neighborhood: r.Body.Consult.Return.Neighborhood,
	}

	return cep, nil
}
