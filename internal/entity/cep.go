package entity

import (
	"github.com/igorhalfeld/lagoinha/pkg/formater"
	"github.com/igorhalfeld/lagoinha/pkg/validator"
)

// Cep standard cep struct
type Cep struct {
	Cep          string `json:"cep"`
	Street       string `json:"street"`
	Neighborhood string `json:"neighborhood"`
	City         string `json:"city"`
	State        string `json:"state"`
	Provider     string `json:"provider"`
}

func (c *Cep) IsValid() bool {
	c.Cep = formater.RemoveSpecialCharacters(c.Cep)

	if c.Cep == "" || validator.ExceedsCepMaximumSize(c.Cep) {
		return false
	}

	c.Cep = formater.LeftPadWithZeros(c.Cep)

	return true
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
