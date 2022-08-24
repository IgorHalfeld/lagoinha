package lagoinha

import (
	"github.com/igorhalfeld/lagoinha/internal/entity"
	"github.com/igorhalfeld/lagoinha/internal/service/brasilapi"
	"github.com/igorhalfeld/lagoinha/internal/service/viacep"
	"github.com/igorhalfeld/lagoinha/pkg/errors"
)

func getAddress(cepRaw string, chResponse chan *entity.Cep, chError chan error) {
	cep := entity.Cep{
		Cep: cepRaw,
	}

	if !cep.IsValid() {
		chError <- errors.CepNotValidError
		return
	}

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

// GetAddress - get address
func GetAddress(cepRaw string) (chan *entity.Cep, chan error) {
	chResponse := make(chan *entity.Cep)
	chError := make(chan error)

	go getAddress(cepRaw, chResponse, chError)

	return chResponse, chError
}
