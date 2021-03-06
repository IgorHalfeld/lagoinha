package lagoinha

import (
	"reflect"

	"github.com/igorhalfeld/lagoinha/services"
	"github.com/igorhalfeld/lagoinha/structs"
	"github.com/igorhalfeld/lagoinha/utils"
)

// GetAddress - get address
func GetAddress(cep string) (*structs.Cep, error) {
	services := services.Container{
		CorreiosService: services.NewCorreiosService(),
		ViaCepService:   services.NewViaCepService(),
		WidenetService:  services.NewWidenetService(),
	}

	cepValidated := utils.RemoveSpecialCharacters(cep)
	cepValidated = utils.LeftPadWithZeros(cep)

	respCh := make(chan *structs.Cep)
	errCh := make(chan error)

	var servicesCount int = reflect.TypeOf(services).NumField()
	var errorsCount []error

	// @TODO: If services be more than 3,
	// dispatch this goroutines dynamically
	go func(cv string) {
		c, err := services.CorreiosService.Request(cv)
		if err != nil {
			errorsCount = append(errorsCount, err)
			if len(errorsCount) == servicesCount {
				errCh <- err
				respCh <- nil
			}
		}
		if c != nil {
			respCh <- c
			errCh <- nil
		}
	}(cepValidated)

	go func(cv string) {
		c, err := services.ViaCepService.Request(cv)
		if err != nil {
			errorsCount = append(errorsCount, err)
			if len(errorsCount) == servicesCount {
				errCh <- err
				respCh <- nil
			}
		}
		if c != nil {
			respCh <- c
			errCh <- nil
		}
	}(cepValidated)

	go func(cv string) {
		c, err := services.WidenetService.Request(cv)
		if err != nil {
			errorsCount = append(errorsCount, err)
			if len(errorsCount) == servicesCount {
				errCh <- err
				respCh <- nil
			}
		}
		if c != nil {
			respCh <- c
			errCh <- nil
		}
	}(cepValidated)

	return <-respCh, <-errCh
}
