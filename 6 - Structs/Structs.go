package main

import "fmt"

type usuario struct {
	nome  string
	idade int8
}

type endereco struct {
	logadouro string
	numero    int8
}

func main() {

	fmt.Println("Estrutura do Structs")

	var u usuario // Aqui criamos uma estrutura e mostramos algumas formas de declarar.
	u.nome = "Thiago"
	u.idade = 35
	fmt.Println(u)

	enderecoexemplo := endereco{"Rua dos bobos", 0}

	usuario2 := usuario{"Thiago", 35} //Segunda forma de declarar.
	fmt.Println(usuario2, "E o endereço é", enderecoexemplo)

	usuario3 := usuario{idade: 22}
	fmt.Println(usuario3)
}
