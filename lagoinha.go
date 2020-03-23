package lagoinha

import (
	"github.com/igorhalfeld/lagoinha/containers"
	"github.com/igorhalfeld/lagoinha/handlers"
	"github.com/igorhalfeld/lagoinha/services"
	"github.com/igorhalfeld/lagoinha/structs"
)

// GetAddress - get address
func GetAddress(cep string) (*structs.Cep, error) {

	con := containers.Container{
		CorreiosService: services.NewCorreiosService(),
		ViaCepService:   services.NewViaCepService(),
	}

	return handlers.
		NewGetAddress(con).
		Run(cep)
}
