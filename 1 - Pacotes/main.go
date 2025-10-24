package main

import (
	"fmt"
	"modulo/auxiliar"

	"github.com/badoux/checkmail"
)

func main() {
	fmt.Println("Escrevendo do arquivo main")
	auxiliar.Escrever()
	auxiliar.Escrever2()

	erro := checkmail.ValidateFormat("gabrielvagula@gmail.com")
	fmt.Println(erro)
}
