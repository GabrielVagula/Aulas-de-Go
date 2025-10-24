package main

import "fmt"

func Generica(interf interface{}) {
	fmt.Println(interf)
}

func main() {
	Generica("String")
	Generica(1)
	Generica(true)

	fmt.Println("Stringa", 2, 4, 5, true, float64(1231241254363))

}
