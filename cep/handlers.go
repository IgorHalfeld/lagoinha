package cep

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/utils"
)

// Cep - get address
func Cep(cep string) (interface{}, error) {
	cepValidated := utils.RemoveSpecialCharacters(cep)
	if utils.ValidateInputLength(cepValidated) == false {
		return nil, fmt.Errorf("Cep length exceeds maximum allowed")
	}
	cepValidated = utils.LeftPadWithZeros(cep)

	address, err := utils.RaceServices(cepValidated)
	return address, err
}
