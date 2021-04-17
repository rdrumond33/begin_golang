package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var contador int

func main() {
	quantidadeRoutnes := 100

	incrementar(quantidadeRoutnes)
	wg.Wait()
	fmt.Println("O valor do contador e ", contador)
}

func incrementar(goroutnes int) {
	wg.Add(goroutnes)

	for i := 0; i < goroutnes; i++ {
		go func() {
			defer wg.Done()

			a := contador
			fmt.Println("O valor do contador e ", a)
			runtime.Gosched()
			a++
			contador = a
		}()
	}
}
