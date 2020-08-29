package services

import "github.com/igorhalfeld/lagoinha/structs"

// ServiceInterface maintains integrity of service methods
type ServiceInterface interface {
	Request(cep string) (*structs.Cep, error)
}

// Container for services
type Container struct {
	CorreiosService ServiceInterface
	ViaCepService   ServiceInterface
}
