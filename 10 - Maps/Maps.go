package main

import (
	"fmt"
)

func main() {

	usuario := map[string]string{ // Não se pode trocar o tipo, pq ele não permite, apenas um numero nesse caso.
		// Alterando o primeiro paramentro (chave) ele me pede um numero e depois um string.
		"nome":    "Pedro",
		"sobreno": "Silva",
	} // Dentro do colchete "[]" é o tipo da chave e fora dele é o tipo dos valores
	fmt.Println(usuario["nome"]) // Para chamar um  dado escolhido, so colocar entre conchetes e aspas [] e "".

	// **MANTER OS TIPOS COERENTES

	// Aninhamento de map.

	usuario2 := map[string]map[string]string{ // Um map dentro de outro map passando valores
		"nome": {
			"Primeiro":  "Gabriel",
			"Sobrenome": "Vagula Gomes",
		}, //Precisa colocar uma virgula no final do map.
		"curso": {
			"Nome":      "Tecnologia",
			"Faculdade": "FATEC",
		},
	}

	fmt.Println(usuario2)

	delete(usuario2, "nome") // Para deletar uma chave do seu map.

	fmt.Println(usuario2)

	usuario2["carreira"] = map[string]string{
		"nome": "Engenheiro",
	}

	fmt.Println(usuario2)
}
