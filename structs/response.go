package structs

// Response standard response struct
type Response struct {
	Data Cep
	Err  error
}

// ViaCepResponse Via cep response
type ViaCepResponse struct {
	Cep          string `json:"cep"`
	State        string `json:"uf"`
	City         string `json:"localidade"`
	Neighborhood string `json:"bairro"`
	Street       string `json:"logradouro"`
}

// CorreiosResponse Correios response
type CorreiosResponse struct {
	Body struct {
		Consult correiosBody `xml:"consultaCEPResponse"`
	} `xml:"Body"`
}

type correiosBody struct {
	Return struct {
		Cep          string `xml:"cep"`
		State        string `xml:"uf"`
		City         string `xml:"cidade"`
		Neighborhood string `xml:"bairro"`
		Street       string `xml:"end"`
	} `xml:"return"`
}
