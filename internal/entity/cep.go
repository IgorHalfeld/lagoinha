package entity

import "github.com/igorhalfeld/lagoinha/pkg/validator"

// Cep standard cep struct
type Cep struct {
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Provider     string `json:"provider"`
}

func (c *Cep) ApplyFormaterAndLinters() {
	c.Cep = validator.RemoveSpecialCharacters(c.Cep)
	c.Cep = validator.LeftPadWithZeros(c.Cep)
}

func (c *Cep) IsValid() bool {
	return validator.ValidateInputLength(c.Cep)
}

func (c *Cep) HasAllAddressInfo() bool {
	if c.Cep != "" &&
		c.Street != "" &&
		c.Neighborhood != "" &&
		c.City != "" &&
		c.State != "" &&
		c.Provider != "" {
		return true
	}

	return false
}
