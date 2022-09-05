package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	fmt.Println("Total amount of cep providers:", lagoinha.GetTotalAmountOfCepProviders())
	chResp, chErr := lagoinha.GetAddress("04568000", &lagoinha.GetAddressOptions{
		PreferenceForAPI: "ViaCEP",
	})

	select {
	case address := <-chResp:
		fmt.Printf("Response: %+v\n", address)
	case err := <-chErr:
		fmt.Printf("Error: %+v\n", err)
	}
}
