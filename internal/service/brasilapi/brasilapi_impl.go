package brasilapi

import (
	"encoding/json"
	"net/http"

	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

type BrasilAPIService struct {
}

func New() *BrasilAPIService {
	return &BrasilAPIService{}
}

func (ba *BrasilAPIService) Request(cep string) (*entity.Cep, error) {
	result := brasilAPIResponse{}

	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
	if err != nil {
		return nil, err
	}

	defer res.Body.Close()

	err = json.NewDecoder(res.Body).Decode(&result)
	if err != nil {
		return nil, err
	}

	return ba.formater(&result)
}

func (ba *BrasilAPIService) formater(r *brasilAPIResponse) (*entity.Cep, error) {
	if r == nil {
		return nil, errors.CepNotFoundError
	}

	cep := &entity.Cep{
		Cep:          r.Cep,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		State:        r.State,
		Street:       r.Street,
		Provider:     "BrasilAPI",
	}

	return cep, nil
}
