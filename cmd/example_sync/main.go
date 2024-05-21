package main

import (
	"fmt"
	"log"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	fmt.Println("Total amount of cep providers:", lagoinha.GetTotalAmountOfCepProviders())
	address, err := lagoinha.GetAddressSync("15809240", &lagoinha.GetAddressOptions{
		PreferenceForAPI: "Apicep",
	})

	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	fmt.Printf("Response: %+v\n", address)
}
