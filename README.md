<p align="center">
  <img src="assets/logo.png" width="100px" />
  <h3 align="center">
    Lagoinha
  </h3>
  <p align="center">
    Utilitário Golang para busca por CEP integrado diretamente <br /> aos serviços dos Correios, ViaCEP e outros
  </p>
</p>

--- 
<a href='https://github.com/jpoles1/gopherbadger' target='_blank'>![gopherbadger-tag-do-not-edit](https://img.shields.io/badge/Go%20Coverage-16%25-brightgreen.svg?longCache=true&style=flat)</a>

Lagoinha é basicamente um pacote que usa a API dos Correios, ViaCep e outros para pegar o endereço com base em um CEP. O que o pacote faz, é disparar pra todas as APIs ao mesmo tempo e retornar com o resultado da primeira API que responder.

### Por que esse nome

É simples, veja o [vídeo](https://www.youtube.com/watch?v=C1Sd_RWF5ks)!
(onde é que eu tô, lagoinha, CEP, endereço...)

### Instalação

```sh
go get -u https://github.com/IgorHalfeld/lagoinha.git
```

### Como usar

```golang
package main

import (
	"fmt"

	"github.com/igorhalfeld/lagoinha"
)

func main() {
	chResp, chErr := lagoinha.GetAddress("04568000")

	select {
	case address := <-chResp:
		fmt.Printf("Response: %+v\n", address)
	case err := <-chErr:
		fmt.Printf("Error: %+v\n", err)
	}
}
```

logo by [@nelsonsecco](https://twitter.com/nelsonsecco)
