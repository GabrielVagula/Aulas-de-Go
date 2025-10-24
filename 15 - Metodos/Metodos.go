package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func (u usuario) salvar() {
	fmt.Printf("Salvando os dados do usuario %s no banco de dados\n", u.nome)
}

func (u usuario) MaiordeIdade() bool {
	return u.idade >= 18
}

func (u *usuario) FelizAniversario() {
	u.idade++
}

func main() {
	Usuario1 := usuario{"Usuario1", 20}
	Usuario1.salvar()

	Usuario2 := usuario{"Kaio", 16}
	Usuario2.salvar()

	MaiordeIdade := Usuario1.MaiordeIdade()
	fmt.Println(MaiordeIdade)

	Usuario1.FelizAniversario()
	fmt.Println(Usuario1.idade)
}
