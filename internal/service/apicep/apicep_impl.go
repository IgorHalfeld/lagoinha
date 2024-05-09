package apicep

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

type ApicepService struct{}

func New() *ApicepService {
	return &ApicepService{}
}

// Request - fetch data from viacep api
func (wn *ApicepService) Request(cep string) (*entity.Cep, error) {
	result := apicepResponse{}

	client := &http.Client{
		Timeout: time.Second * 2,
	}

	res, err := client.Get("https://ws.apicep.com/cep/" + cep + ".json")
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

	if result.Status != 200 {
		return nil, errors.HttpError(result.Status)
	}
	return wn.formater(&result)
}

func (wn *ApicepService) formater(r *apicepResponse) (*entity.Cep, error) {
	if r == nil {
		return nil, errors.CepNotFoundError
	}

	cep := &entity.Cep{
		Cep:          r.Code,
		City:         r.City,
		Neighborhood: r.District,
		State:        r.State,
		Street:       r.Address,
		Provider:     "Apicep",
	}

	return cep, nil
}
