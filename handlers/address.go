package handlers

import (
	"errors"

	"github.com/igorhalfeld/lagoinha/containers"
	"github.com/igorhalfeld/lagoinha/interfaces"
	"github.com/igorhalfeld/lagoinha/structs"
	"github.com/igorhalfeld/lagoinha/utils"
)

// GetAddress handles load address
type GetAddress struct {
	ViaCepService   interfaces.IService
	CorreiosService interfaces.IService
}

// NewGetAddress creates a new instance
func NewGetAddress(container containers.Container) GetAddress {
	return GetAddress{
		CorreiosService: container.CorreiosService,
		ViaCepService:   container.ViaCepService,
	}
}

// Run get address for a given cep
func (ga GetAddress) Run(cep string) (*structs.Cep, error) {
	cepValidated := utils.RemoveSpecialCharacters(cep)
	if utils.ValidateInputLength(cepValidated) == false {
		return nil, errors.New("Cep length exceeds maximum allowed")
	}
	cepValidated = utils.LeftPadWithZeros(cep)

	// var address structs.Cep
	// var err error

	// example
	c, err := ga.CorreiosService.Request(cep)

	return c, err
}
