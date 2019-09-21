package lagoinha

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/utils"
)

// GetAddress - get address
func GetAddress(cep string) (interface{}, error) {
	cepValidated := utils.RemoveSpecialCharacters(cep)
	if utils.ValidateInputLength(cepValidated) == false {
		return nil, fmt.Errorf("Cep length exceeds maximum allowed")
	}
	cepValidated = utils.LeftPadWithZeros(cep)

	address, err := utils.RaceServices(cepValidated)
	return address, err
}
