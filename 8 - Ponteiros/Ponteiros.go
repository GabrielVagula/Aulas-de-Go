package main

import "fmt"

func main() {
	fmt.Println("Ponteiros")

	var variavel1 int8 = 10
	var variavel2 = variavel1
	fmt.Println(variavel1, variavel2)

	variavel1++
	fmt.Println(variavel1, variavel2)

	// Ponteiro e uma referencia de memoria.

	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3 // No print da uma loucura de dado, que é aonde o valor esta colocado na minha memoria do pc.
	// Coloca o "&" Para ele atribuir esse endereço de memoria
	fmt.Println(variavel3, ponteiro) // Aqui o ponteiro passa somente o endereço de memoria

	variavel3 = 150
	fmt.Println(variavel3, *ponteiro) // Quando se coloca o "*" ele mostra oq esta no endereço de memoria.

}
