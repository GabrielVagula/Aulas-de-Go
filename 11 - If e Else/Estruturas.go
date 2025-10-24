package main

import "fmt"

func main() {

	numero := -8

	if numero > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor que 15")
	}

	// Variavel IfNit. Aqui nos declaramos a variavel dentro do proprio if,
	// demos o valor de numero (variavel criada la encima) e ja passamos a condição dentro da estrutura.

	// Essa variavel so serve para esse if, se tentarmos utilziar em outra coisa fora, o GO limita a esse escopo.
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("Maior que zero")
	} else if numero < -10 {
		fmt.Println("Menor que -10")
	} else {
		fmt.Println("O numero esta entre -10 e 0")
	}
}
