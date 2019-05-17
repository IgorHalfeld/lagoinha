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

	"github.com/igorhalfeld/lagoinha/cep"
)

func main() {
	address, _ := cep.Cep("CEP_GOES_HERE")
	fmt.Printf("Complete Address %v:", address)
}
```

logo by [@nelsonsecco](https://twitter.com/nelsonsecco)