package main

import (
	"errors"
	"fmt"
)

func main() {
	//Numeros Reais
	//int8, int16, int 32, int 64
	var numero int32 = 100000000
	fmt.Println(numero)

	var numero1 uint32 = 10000
	fmt.Println(numero1)

	var numero2 rune = 12345 //igual o int32
	fmt.Println(numero2)

	var numero3 byte = 123 //igual o int8
	fmt.Println(numero3)

	var numeroreal1 float32 = 123.45
	fmt.Println(numeroreal1)

	var numeroreal2 float32 = 1230000.45
	fmt.Println(numeroreal2)

	//também funciona com a inferencia ":=" e segue a arquitetura, se o seu for 32 bits ele gera um float no max 32.
	//Fim Numeros Reais

	//Strings
	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto1"
	fmt.Println(str2)

	char := 'H' //Mostra o valor correspondete a tabela ask
	fmt.Println(char)
	//Fim Strings

	var boleano bool = true //Se deixar sem declarar, ele fica como falso.
	fmt.Println(boleano)

	var erro error = errors.New("Erro Interno!")
	fmt.Println(erro)
}
