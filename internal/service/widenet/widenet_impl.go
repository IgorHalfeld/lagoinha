package widenet

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

type WidenetService struct{}

func New() *WidenetService {
	return &WidenetService{}
}

// Request - fetch data from viacep api
func (wn *WidenetService) Request(cep string) (*entity.Cep, error) {
	result := widenetResponse{}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, err := client.Get("http://apps.widenet.com.br/busca-cep/api/cep/" + cep + ".json")
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

	return wn.formater(&result)
}

func (wn *WidenetService) formater(r *widenetResponse) (*entity.Cep, error) {
	if r == nil {
		return nil, errors.CepNotFoundError
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
