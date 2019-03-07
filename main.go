package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/cep"
)

func main() {
	address, err := cep.Cep("CEP_GOES_HERE")
	fmt.Printf("Complete Address %v:", address)
	fmt.Printf("\nError %v:", err)
}
