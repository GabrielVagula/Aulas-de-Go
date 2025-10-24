package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct { //Não da erro pq uma estrutura pode ser feita sem ser usada
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("Herança")

	p1 := pessoa{"Vanessa", "Vagula", 24, 165}
	fmt.Println(p1)

	e1 := estudante{p1, "Engenharia", "USP"}
	fmt.Println(e1.nome) // Depois do ponto voce pode pegar oq quiser colocando um ponto e puxando o atributo

}

// Como se fosse uma pseudo herança
