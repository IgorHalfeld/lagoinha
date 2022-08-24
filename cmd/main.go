package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	chResp, chErr := lagoinha.GetAddress("04568000", nil)

	select {
	case address := <-chResp:
		fmt.Printf("Response: %+v\n", address)
	case err := <-chErr:
		fmt.Printf("Error: %+v\n", err)
	}
}
