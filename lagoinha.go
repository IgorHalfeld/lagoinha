package lagoinha

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/internal/service"
	"github.com/igorhalfeld/lagoinha/internal/service/apicep"
	"github.com/igorhalfeld/lagoinha/internal/service/brasilapi"
	"github.com/igorhalfeld/lagoinha/internal/service/viacep"
	"github.com/igorhalfeld/lagoinha/internal/service/widenet"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

var providers = map[string]service.Provider{
	"BrasilAPI": brasilapi.New(),
	"ViaCEP":    viacep.New(),
	"Apicep":    apicep.New(),
	"WidNet":    widenet.New(),
}

// GetTotalAmountOfCepProviders returns amount of current enabled cep provivers
func GetTotalAmountOfCepProviders() int {
	return len(providers)
}

func getAddress(cepRaw string, opts *GetAddressOptions, chResponse chan *entity.Cep, chError chan error) {
	var wg sync.WaitGroup

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
			chResponseInternal := make(chan *entity.Cep, GetTotalAmountOfCepProviders())
			chErrorInternal := make(chan error)

			provider, ok := providers[opts.PreferenceForAPI]
			if !ok {
				chError <- errors.PreferenceProviderNotFound
				return
			}

			ctx, cancel := context.WithTimeout(context.Background(), 500*time.Millisecond)
			defer cancel()

			wg.Add(1)
			go func() {
				defer wg.Done()
				c, err := provider.Request(cep.Cep)
				if err != nil {
					chErrorInternal <- err
					return
				}

				chResponseInternal <- c
			}()

			select {
			case address := <-chResponseInternal:
				chResponse <- address
			case err := <-chErrorInternal:
				chError <- err
			case <-ctx.Done():
				cancel()
			}
		}
	}

	// TODO: add context.WithCancel for slower requests
	for providerName := range providers {
		wg.Add(1)
		go func(p, cv string) {
			defer wg.Done()
			c, err := providers[p].Request(cv)
			if err != nil {
				chError <- err
				return
			}

			chResponse <- c
		}(providerName, cep.Cep)
	}

	wg.Wait()
	close(chResponse)
	close(chError)
}

type GetAddressOptions struct {
	PreferenceForAPI string
}

// GetAddress - get address
func GetAddress(cepRaw string, opts *GetAddressOptions) (chan *entity.Cep, chan error) {
	chResponse := make(chan *entity.Cep, GetTotalAmountOfCepProviders())
	chError := make(chan error)

	go getAddress(cepRaw, opts, chResponse, chError)

	return chResponse, chError
}

func GetAddressSync(cepRaw string, opts *GetAddressOptions) (*entity.Cep, error) {
	totalAmountOfCepProvider := GetTotalAmountOfCepProviders()
	chResponse := make(chan *entity.Cep, totalAmountOfCepProvider)
	chError := make(chan error)

	go getAddress(cepRaw, opts, chResponse, chError)

	var lastError error
	for i := 0; i <= totalAmountOfCepProvider; i++ {
		select {
		case address := <-chResponse:
			return address, nil
		case err := <-chError:
			lastError = err
		}
	}

	if lastError != nil {
		return nil, lastError
	}

	return nil, fmt.Errorf("address not found for cep %s", cepRaw)
}
