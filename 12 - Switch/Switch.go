package main

import "fmt"

// Aqui voce usa o switch de numero int q devolve uma string. Parecido com o if/else
func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda Feira"
	case 3:
		return "Terça Feira"
	case 4:
		return "Quarta Feira"
	case 5:
		return "Quinta Feira"
	case 6:
		return "Sexta Feira"
	case 7:
		return "Domingo"
	default:
		return "Numero invalido"
	}
}

// Aqui uma outra forma de usar o switch
func diaDaSemana2(numero int) string {
	var diaDaSemana string

	switch {
	case numero == 1:
		diaDaSemana = "Domingo"
		fallthrough // Ele joga para a proxima condição. Só para saber que existe.
	case numero == 2:
		diaDaSemana = "Segunda"
	case numero == 3:
		diaDaSemana = "Terça"
	case numero == 4:
		diaDaSemana = "Quarta"
	case numero == 5:
		diaDaSemana = "Quinta"
	case numero == 6:
		diaDaSemana = "Sexta"
	case numero == 7:
		diaDaSemana = "Sabado"
	default:
		diaDaSemana = "Numero Invalido"
	}

	return diaDaSemana
}

func main() {

	dia := diaDaSemana(100)
	fmt.Println(dia)

	dia2 := diaDaSemana2(10)
	fmt.Println(dia2)

}
