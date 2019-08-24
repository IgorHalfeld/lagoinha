package utils_test

import (
	"testing"

	"github.com/igorhalfeld/lagoinha/utils"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCEP(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "CEP Suite")
}

var _ = Describe("CEP", func() {
	var validCepWithSpecialCharacters = "01310-200"
	var validCepWithoutSpecialCharacters = "013102000"
	var invalidCepWithoutSpecialCharacters = "01310200"
	var cepThatNeedsBePadedWithZeros = "013102"

	Describe("All helper functions should work correctly", func() {
		It("Should remove special characters", func() {
			Expect(utils.RemoveSpecialCharacters(validCepWithSpecialCharacters)).To(Equal("01310200"))
		})

		It("Should validate input length", func() {
			Expect(utils.ValidateInputLength(invalidCepWithoutSpecialCharacters)).To(Equal(true))
			Expect(utils.ValidateInputLength(validCepWithoutSpecialCharacters)).To(Equal(false))
		})

		It("Should pad left with zeros", func() {
			Expect(utils.LeftPadWithZeros(cepThatNeedsBePadedWithZeros)).To(Equal("00013102"))
		})
	})
})
