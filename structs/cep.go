package structs

// Cep standard cep struct
type Cep struct {
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Provider     string `json:"provider"`
}
