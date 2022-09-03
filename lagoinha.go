package lagoinha

import (
	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/internal/service"
	"github.com/igorhalfeld/lagoinha/internal/service/brasilapi"
	"github.com/igorhalfeld/lagoinha/internal/service/viacep"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

var providers = map[string]service.Provider{
	"BrasilAPI": brasilapi.New(),
	"ViaCEP":    viacep.New(),
}

func GetTotalAmountOfCepProviders() int {
	return len(providers)
}

// TotalAmountOfCepProviders returns amount of current enabled cep provivers
const TotalAmountOfCepProviders = 2

func getAddress(cepRaw string, opts *GetAddressOptions, chResponse chan *entity.Cep, chError chan error) {
	cep := entity.Cep{
		Cep: cepRaw,
	}

	if !cep.IsValid() {
		chError <- errors.CepNotValidError
		return
	}

	// TODO: enhance this way of handling options
	if opts != nil {
		if opts.PreferenceForAPI != "" {
			provider, ok := providers[opts.PreferenceForAPI]
			if !ok {
				chError <- errors.PreferenceProviderNotFound
			}

			c, err := provider.Request(cep.Cep)
			if err != nil {
				chError <- err
				return
			}

			chResponse <- c
			return
		}
	}

	// TODO: add context.WithCancel for slower requests
	for providerName := range providers {
		go func(p, cv string) {
			c, err := providers[p].Request(cv)
			if err != nil {
				chError <- err
				return
			}

			chResponse <- c
		}(providerName, cep.Cep)
	}
}

type GetAddressOptions struct {
	PreferenceForAPI string
}

// GetAddress - get address
func GetAddress(cepRaw string, opts *GetAddressOptions) (chan *entity.Cep, chan error) {
	chResponse := make(chan *entity.Cep, TotalAmountOfCepProviders)
	chError := make(chan error)

	go getAddress(cepRaw, opts, chResponse, chError)

	return chResponse, chError
}
