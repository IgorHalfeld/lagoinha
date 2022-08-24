package lagoinha

import (
	"github.com/igorhalfeld/lagoinha/internal/entity"
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

	chResponse <- &cep
}

// GetAddress - get address
func GetAddress(cepRaw string) (chan *entity.Cep, chan error) {
	chResponse := make(chan *entity.Cep)
	chError := make(chan error)

	go getAddress(cepRaw, chResponse, chError)

	return chResponse, chError
}
