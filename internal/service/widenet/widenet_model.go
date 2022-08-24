package widenet

type widenetResponse struct {
	Cep          string `json:"code"`
	State        string `json:"state"`
	City         string `json:"city"`
	Neighborhood string `json:"district"`
	Street       string `json:"address"`
}
