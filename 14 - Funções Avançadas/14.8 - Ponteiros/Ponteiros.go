package main

import "fmt"

func InverteroSinal(numero int) int {
	return numero * -1
}

func InverteoSinalComPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 20
	numeroInvertido := InverteroSinal(numero)
	fmt.Println(numeroInvertido)
	fmt.Println(numero)

	novonumero := 40
	fmt.Println(novonumero)
	InverteoSinalComPonteiro(&novonumero)
	fmt.Println(novonumero)
}
