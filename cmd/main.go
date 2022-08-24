package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	chResp, chErr := lagoinha.GetAddress("")

	select {
	case address := <-chResp:
		fmt.Printf("Response: %+v\n", address)
	case err := <-chErr:
		fmt.Printf("Response: %+v\n", err)
	}
}