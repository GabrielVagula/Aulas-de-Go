package main

import "fmt"

func main() {

	retorno := func(texto string) string {
		return fmt.Sprintf("Recebido -> %s %d", texto, 10)
	}("Passando parametro")

	fmt.Println(retorno)

}

// Aqui utilizamos a função anonima, que representa o () depois da função.
// Tambem utilizamos o %s que serve para concatenar com o texto da função.
// Tambem utilizamos o %d para concatenar os numeros.
