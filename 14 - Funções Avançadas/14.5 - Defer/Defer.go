package main

import "fmt"

func funcao1() {
	fmt.Println("Executando a função 1")
}

func funcao2() {
	fmt.Println("Executando a função 2")
}

func alunoAprovado(n1, n2 float32) bool {
	defer fmt.Println("Media das notas calculadas. Resultado sera retornado")
	fmt.Println("Função para ver se o aluno esta aprovado")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	defer funcao1() // Aqui estamos adiando com o defer a função, ele executa antes da função main terminar.
	funcao2()

	fmt.Println(alunoAprovado(7, 8))

}
