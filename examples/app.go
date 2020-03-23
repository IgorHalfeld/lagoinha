package example

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	address, err := lagoinha.GetAddress("01310200")
	fmt.Printf("Complete Address %v:", address)
	fmt.Printf("\nError %v:", err)
}
