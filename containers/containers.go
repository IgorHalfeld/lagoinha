package containers

import "github.com/igorhalfeld/lagoinha/interfaces"

// Container is a basic container
type Container struct {
	CorreiosService interfaces.IService
	ViaCepService   interfaces.IService
}
