package services_test

import (
	"testing"

	"github.com/igorhalfeld/lagoinha/services"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCepAbertoService(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CepAberto Service Suite")
}

var _ = Describe("", func() {
	// var response = make(chan models.Status)

	It("Should return corretly address with cep provied", func() {
		Expect(services.FetchCepAbertoService)
	})
})
