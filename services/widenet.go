package services

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/structs"
)

// ViaCepService service
type WidenetService struct{}

// NewViaCepService creates a new instance
func NewWidenetService() *WidenetService {
	return &WidenetService{}
}

// Request - fetch data from viacep api
func (wn *WidenetService) Request(cep string) (*structs.Cep, error) {
	result := structs.WidenetResponse{}

	res, err := http.Get("http://apps.widenet.com.br/busca-cep/api/cep/" + cep + ".json")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return wn.formater(&result)
}

func (wn *WidenetService) formater(r *structs.WidenetResponse) (*structs.Cep, error) {
	if r == nil {
		return nil, errors.New("Cep not found")
	}

	cep := &structs.Cep{
		Cep:          r.Cep,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		State:        r.State,
		Street:       r.Street,
		Provider:     "Widenet",
	}

	return cep, nil
}
