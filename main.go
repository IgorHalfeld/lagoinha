package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/cep"
)

func main() {
	address, err := cep.Cep("01310200")
	fmt.Printf("Complete Address %v:", address)
	fmt.Printf("\nError %v:", err)
}
