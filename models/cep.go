package models

// Status - service status response
type Status struct {
	Ok    bool
	Value interface{}
}

// ViaCepResponse - Via cep response
type ViaCepResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

// CepAbertoResponse - Cepaberto Response
type CepAbertoResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"estado"`
	City         string `json:"cidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

// CorreiosResponse - Correios response
type CorreiosResponse struct {
	Body struct {
		Consult completeResponse `xml:"consultaCEPResponse"`
	} `xml:"Body"`
}

type completeResponse struct {
	Return struct {
		Cep          string `xml:"cep"`
		State        string `xml:"uf"`
		City         string `xml:"cidade"`
		Neighborhood string `xml:"bairro"`
		Street       string `xml:"end"`
	} `xml:"return"`
}
