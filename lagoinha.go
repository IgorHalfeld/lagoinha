package lagoinha

import (
	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/internal/service/brasilapi"
	"github.com/igorhalfeld/lagoinha/internal/service/viacep"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

// TotalAmountOfCepProviders returns amount of current enabled cep provivers
const TotalAmountOfCepProviders = 2

func getAddress(cepRaw string, chResponse chan *entity.Cep, chError chan error) {
	cep := entity.Cep{
		Cep: cepRaw,
	}

	if !cep.IsValid() {
		chError <- errors.CepNotValidError
		return
	}

	// TODO: add context.WithCancel for slower requests
	go func(cv string) {
		service := viacep.New()
		c, err := service.Request(cv)
		if err != nil {
			chError <- err
			return
		}

		chResponse <- c
	}(cep.Cep)

	go func(cv string) {
		service := brasilapi.New()
		c, err := service.Request(cv)
		if err != nil {
			chError <- err
			return
		}

		chResponse <- c
	}(cep.Cep)
}

type GetAddressOptions struct {
	PreferenceForAPI string
}

// GetAddress - get address
func GetAddress(cepRaw string, _ *GetAddressOptions) (chan *entity.Cep, chan error) {
	chResponse := make(chan *entity.Cep, TotalAmountOfCepProviders)
	chError := make(chan error)

	go getAddress(cepRaw, chResponse, chError)

	return chResponse, chError
}
