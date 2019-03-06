package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/cep"
)

func main() {
	address, _ := cep.Cep("01307-000")
	fmt.Printf("Complete Address %v:", address)
}
