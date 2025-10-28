// Teste de unidade

package enderecos

import "testing"

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipoDeEnderecos(t *testing.T) {

	cenarioDeTeste := []cenarioDeTeste{
		{"Viela KKKKK", "Tipo Invalido"},
		{"Rua 123", "Rua"},
		{"Avenida Porto", "Avenida"},
		{"Rodovia Kill", "Rodovia"},
		{"Rua Abras", "Rua"},
		{"Praça 123", "Tipo Invalido"},
		{"Avenida Paulista", "Avenida"},
		{"Beco Diagonal", "Tipo Invalido"},
	}

	for _, cenario := range cenarioDeTeste {
		tipoDeEnderecoRecebido := TipoDeEnderecos(cenario.enderecoInserido)
		if tipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				tipoDeEnderecoRecebido,
				cenario.retornoEsperado,
			)
		}
	}

	//if tipoDeEnderecoRecebido != tipoDeEnderecoEsperado {
	//	t.Errorf("O tipo recebido é diferente do tipo esperado! Esperava o tipo %s e recebi o tipo %s !",
	//		tipoDeEnderecoEsperado,
	//		tipoDeEnderecoRecebido)
	//}

}
