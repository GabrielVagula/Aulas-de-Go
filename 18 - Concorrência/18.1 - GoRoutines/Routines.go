package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Ola mundo")
	escrever("Outro Ola mundo")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
