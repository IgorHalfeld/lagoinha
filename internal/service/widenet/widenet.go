package widenet

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/igorhalfeld/lagoinha/internal/entity"
)

type WidenetService struct{}

func New() *WidenetService {
	return &WidenetService{}
}

// Request - fetch data from viacep api
func (wn *WidenetService) Request(cep string) (*entity.Cep, error) {
	result := widenetResponse{}

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

func (wn *WidenetService) formater(r *widenetResponse) (*entity.Cep, error) {
	if r == nil {
		return nil, errors.New("Cep not found")
	}

	cep := &entity.Cep{
		Cep:          r.Cep,
		City:         r.City,
		Neighborhood: r.Neighborhood,
		State:        r.State,
		Street:       r.Street,
		Provider:     "Widenet",
	}

	return cep, nil
}
