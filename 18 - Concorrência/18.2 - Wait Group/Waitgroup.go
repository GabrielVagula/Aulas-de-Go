package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var waitGroup sync.WaitGroup

	waitGroup.Add(4)

	go func() {
		escrever("Ola mundo")
		waitGroup.Done() //Esse .Done faz aquele contador ficar com -1
	}()

	go func() {
		escrever("Outro Ola mundo")
		waitGroup.Done()
	}()

	go func() {
		escrever("Vai")
		waitGroup.Done()
	}()

	go func() {
		escrever("Go")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
