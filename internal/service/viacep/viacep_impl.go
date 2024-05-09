package viacep

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

type ViaCepService struct {
}

// NewViaCepService creates a new instance
func New() *ViaCepService {
	return &ViaCepService{}
}

// Request - fetch data from viacep api
func (vc *ViaCepService) Request(cep string) (*entity.Cep, error) {
	result := viaCepResponse{}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, err := client.Get("https://viacep.com.br/ws/" + cep + "/json/")
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	if res.StatusCode != 200 {
		return nil, errors.HttpError(res.StatusCode)
	}

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return vc.formater(&result)
}

func (vc *ViaCepService) formater(r *viaCepResponse) (*entity.Cep, error) {
	if r == nil {
		return nil, errors.CepNotFoundError
	}

	cep := &entity.Cep{
		Cep:          r.Cep,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		State:        r.State,
		Street:       r.Street,
		Provider:     "ViaCEP",
	}

	return cep, nil
}
