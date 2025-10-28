package main

import (
	"fmt"
	"introducao-testes/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEnderecos("rua Dos Imigrantes")
	fmt.Println(tipoEndereco)

}
