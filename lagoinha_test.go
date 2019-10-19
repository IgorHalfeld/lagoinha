package lagoinha

import (
	"testing"
	"github.com/igorhalfeld/lagoinha/models"
	"github.com/stretchr/testify/assert"
)

func TestGetAddress(t *testing.T) {

	sepInvalidLen := "Hello"
	_, err := GetAddress(sepInvalidLen)
	assert.Error(t, err)

	sepVaild := "01310200"
	address, err := GetAddress(sepVaild)
	expected := models.ViaCepResponse(models.ViaCepResponse{Cep:"01310-200", State:"SP", City:"São Paulo", Neighborhood:"Bela Vista", Street:"Avenida Paulista"})
	assert.Equal(t, expected, address)
	assert.NoError(t, err)

}
