package models

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
