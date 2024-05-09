package main

import (
	"fmt"
	"log"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	fmt.Println("Total amount of cep providers:", lagoinha.GetTotalAmountOfCepProviders())
	address, err := lagoinha.GetAddressWithoutChannel("15809240", &lagoinha.GetAddressOptions{
		PreferenceForAPI: "Apicep",
	})

	if err != nil {
		log.Fatalf("Error: %+v\n", err)
	}

	fmt.Printf("Response: %+v\n", address)
}

// func main() {
// 	fmt.Println("Total amount of cep providers:", lagoinha.GetTotalAmountOfCepProviders())
// 	chResp, chErr := lagoinha.GetAddress("04568000", &lagoinha.GetAddressOptions{
// 		PreferenceForAPI: "Apicep",
// 	})

// 	select {
// 	case address := <-chResp:
// 		fmt.Printf("Response: %+v\n", address)
// 	case err := <-chErr:
// 		fmt.Printf("Error: %+v\n", err)
// 	}
// }
