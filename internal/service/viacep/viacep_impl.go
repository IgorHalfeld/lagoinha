package viacep

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/structs"
)

type ViaCepService struct {
}

// NewViaCepService creates a new instance
func NewViaCepService() ViaCepService {
	return &viaCepImpl{}
}

// Request - fetch data from viacep api
func (vc *ViaCepService) Request(cep string) (*structs.Cep, error) {
	result := viaCepResponse{}

	res, err := http.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return vc.formater(&result)
}

func (vc *ViaCepService) formater(r *structs.ViaCepResponse) (*structs.Cep, error) {
	if r == nil {
		return nil, errors.New("Cep not found")
	}

	cep := &structs.Cep{
		Cep:          r.Cep,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		State:        r.State,
		Street:       r.Street,
		Provider:     "Viacep",
	}

	return cep, nil
}
