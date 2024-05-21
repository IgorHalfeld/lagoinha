<p align="center">
  <img src="assets/logo.png" width="100px" />
  <h3 align="center">
    Lagoinha
  </h3>
  <p align="center">
  	Library to get full address of a Brazilian zip code. <br />
	  Works with VipCep, Correios, and many more.
  </p>
</p>

---

<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-16%25-brightgreen.svg?longCache=true&style=flat)</a>

Lagoinha is a package that uses APIs to find complete addresses by a provided zip code. The lib dispatches several requests at the same time and returns with the one that finished first.

### Por que esse nome

It's a Brazilian meme [vídeo](https://www.youtube.com/watch?v=C1Sd_RWF5ks)!
(onde é que eu tô, lagoinha, CEP, endereço...)

### Install

```sh
go get -u https://github.com/IgorHalfeld/lagoinha.git
```

### How to use

```golang
package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	// get amount of cep providers enabled
	fmt.Println("Total amount of cep providers:", lagoinha.GetTotalAmountOfCepProviders())

	/*
	// if you want to use without handling channels
	addr, err := lagoinha.GetAddressSync("15809240", &lagoinha.GetAddressOptions{
		PreferenceForAPI: "Apicep",
	})
	*/

	chResp, chErr := lagoinha.GetAddress("04568000")

	select {
	case address := <-chResp:
		fmt.Printf("Response: %+v\n", address)
	case err := <-chErr:
		fmt.Printf("Error: %+v\n", err)
	}
}
```

and you can also set a preference api

```golang
chResp, chErr := lagoinha.GetAddress("04568000", &lagoinha.GetAddressOptions{
	PreferenceForAPI: "ViaCEP",
})
```

logo by [@nelsonsecco](https://twitter.com/nelsonsecco)
