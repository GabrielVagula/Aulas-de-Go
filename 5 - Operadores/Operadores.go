package main

import "fmt"

func main() {
	//Aritmeticos
	soma := 1 + 2
	subtração := 5 - 3
	divisao := 10 / 2
	multiplicacao := 2 * 4
	restodadivisao := 10 % 2

	fmt.Println(soma, subtração, divisao, multiplicacao, restodadivisao)

	var numero1 int32 = 10
	var numero2 int32 = 30
	somar := numero1 + numero2
	fmt.Println(somar)
	//Fim Aritmeticos

	//Atribuição
	var variavel1 string = "String1"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	//Fim dos Operadores de Atribuição

	//Operadores Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 != 2)
	//Fim dos Operadores Relacionais

	//Operadores Lógicos
	fmt.Println("-----------------------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso) //AND do Go (todas as condições devem estar em verdadeiro para da verdadeiro)
	fmt.Println(verdadeiro || falso) //OR do Go (uma das condições devem estar em verdadeiro para dar verdadeiro)
	fmt.Println(!verdadeiro)         //Negação do verdadeiro
	fmt.Println(!falso)              //Negação do falso
	//Fim dos Operadores Logicos

	//Operadores Unários
	numero := 10
	numero++
	numero += 15 //Caso eu n queira pular de 1 em 1, igual o exemplo de cima,
	//Eu faço dessa maneira para no exemplo pular de 15 em 15. É igual numero = numero + 15
	fmt.Println(numero)

	numero--
	numero -= 2 //Mesma situação, operador -- e se precisar q subtraia mais, "-=" ajuda com essa situação
	fmt.Println(numero)

	fmt.Println("-----------------------------")
	numeroteste := 10 //Mesma situação, com numeros escolhidos na frente
	numeroteste *= 3
	fmt.Println(numeroteste)
	numeroteste /= 2
	fmt.Println(numeroteste)
	numeroteste %= 3
	fmt.Println(numeroteste)
	//Fim dos Operadores Unários
}
