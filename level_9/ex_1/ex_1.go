package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	fmt.Println("Iniciando")

	go func() {
		defer wg.Done()
		fmt.Println("Hello")
	}()

	go func() {
		defer wg.Done()
		fmt.Println("World!!")
	}()

	wg.Wait()
	fmt.Println("Terminou")
}
