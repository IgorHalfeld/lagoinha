package interfaces

import "github.com/igorhalfeld/lagoinha/structs"

// IService maintains integrity of service methods
type IService interface {
	Request(cep string) (*structs.Cep, error)
}
