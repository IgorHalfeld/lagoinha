package main

import (
	"log"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	address, err := lagoinha.GetAddress("01310200")
	if err != nil {
		log.Fatalf("\nError %v:", err)
	}

	log.Printf("Complete Address %v:", address)
}
