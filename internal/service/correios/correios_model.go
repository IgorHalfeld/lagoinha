package correios

type correiosBody struct {
	Return struct {
		Cep          string `xml:"cep"`
		State        string `xml:"uf"`
		City         string `xml:"cidade"`
		Neighborhood string `xml:"bairro"`
		Street       string `xml:"end"`
	} `xml:"return"`
}

type correiosResponse struct {
	Body struct {
		Consult correiosBody `xml:"consultaCEPResponse"`
	} `xml:"Body"`
}
