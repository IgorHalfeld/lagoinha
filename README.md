<p align="center">
  <h2 align="center">
    Lagoinha
  </h2>
  <p align="center">
    Utilitário Golang para busca por CEP integrado diretamente aos serviços dos Correios, ViaCEP e outros <br />
    <a href="https://www.youtube.com/watch?v=C1Sd_RWF5ks" align="center">
      <img src="assets/lagoinha.png" style="width: 100%" />
    </a>
  </p>
</p>

--- 

### Install

```sh
go get -u https://github.com/IgorHalfeld/lagoinha.git
```

### How to use

```golang
package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha/cep"
)

func main() {
	address, _ := cep.Cep("01307-000")
	fmt.Printf("Complete Address %v:", address)
}
```