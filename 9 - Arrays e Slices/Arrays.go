package main

import "fmt"

func main() {
	fmt.Println("Agora sobre Arrays e Slices")

	var array1 [5]string
	array1[0] = "Ola mundo"
	fmt.Println(array1)

	array2 := [5]string{"Um", "Dois", "Três", "Quatro", "Cinco"} // Os arrays são fixos, se vc tentar colocar mais, ele n permite
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5} // Aqui eu limito o array de acordo com os numero que eu coloquei
	fmt.Println(array3)               // Não deixa com um tamanho dinamo, mas deixa no numero de objetos colocado antes.

	slice := []int{1, 2, 3, 4, 5, 6, 7, 8} // O slice não tem limite de objetos, ele é flexivel com os seus dados.
	fmt.Println(slice)

	slice = append(slice, 9) // Aqui eu acrescento um dado/item no meu slice.
	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição Alterada"
	fmt.Println(slice2)

	//Arrays Internos

	// 3 Parametros, Tipo do slice, Tamanho do slice (quantidade de itens), Capacidade (quantidade maxima de itens)

	fmt.Println("-----------------------")
	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)

	slice3 = append(slice3, 4)
	slice3 = append(slice3, 5)
	fmt.Println(slice3)
	fmt.Println(len(slice3)) // Lenght
	fmt.Println(cap(slice3)) // Capacidade
	// O slice nunca tem um tamanho, porque sempre que vc passa a capacidade ele dobra o valor da capacidade

	slice4 := make([]float32, 5)
	fmt.Println(slice4)
	slice4 = append(slice4, 4)
	fmt.Println(len(slice4)) // Lenght
	fmt.Println(cap(slice4)) // Capacidade
	// o array e uma lista com tamanho fixo e o slice não tem o tamanho fisico.

}
