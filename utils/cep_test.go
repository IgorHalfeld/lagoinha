package utils

import "testing"

func TestRemoveSpecialCharacters(t *testing.T) {
	cep := "01310-200"
	cepWihoutSpecialCharacters := RemoveSpecialCharacters(cep)
	if cepWihoutSpecialCharacters != "01310200" {
		t.Error("Not removed special characters")
	}
}

func TestValidateInputLength(t *testing.T) {
	cepInvalid := "013102000"
	cepTestI := ValidateInputLength(cepInvalid)
	if cepTestI == true {
		t.Error("Not validating cep length correctly")
	}

	cepValid := "01310200"
	cepTestV := ValidateInputLength(cepValid)
	if cepTestV == false {
		t.Error("Not validating cep length correctly")
	}
}

func TestLeftPadWithZeros(t *testing.T) {
	cep := "013102"
	cepWithZeros := LeftPadWithZeros(cep)
	if cepWithZeros != "00013102" {
		t.Error("Error on padding with zeros")
	}
}
