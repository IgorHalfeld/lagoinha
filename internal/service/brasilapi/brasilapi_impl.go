package brasilapi

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/internal/entity"
)

type BrasilAPI struct {
}

func New() BrasilAPI {
	return &BrasilAPI{}
}

func (ba *BrasilAPI) Request(cep string) (*entity.Cep, error) {
	result := BrasilAPIResponse{}

	res, err := http.Get("https://brasilapi.com.br/api/cep/v1/" + cep)
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

func (ba *BrasilAPI) formater(r *BrasilAPIResponse) (*entity.Cep, error) {
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
