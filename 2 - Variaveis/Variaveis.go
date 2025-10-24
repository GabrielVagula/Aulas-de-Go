package main

import (
	"fmt"
)

func main() {
	var variavel1 string = "Exemplo numero 1 - declaração normal de variavel"
	variavel2 := "Exemplo numero 2 - Variavel com inferencia, que é o :=" // := é a inferencia

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 = "Exemplo nuemro 3"
		variavel4 = "criar uma variavel com variaveis dentro dela"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Exemplo numero 4", "Exemplos com ','"
	fmt.Println(variavel5, variavel6)

	const cosntante1 string = "Exemplo numero 5 - declaração de uma constante"
	fmt.Println(cosntante1)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println("Exemplo numero 6 que inverte os valores das variaveis")
	fmt.Println(variavel5, variavel6)
}
