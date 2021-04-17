package main

import (
	"fmt"
	"runtime"
	"sync"
)

var mu sync.Mutex
var wg sync.WaitGroup
var contador int

func main() {
	quantidadeRoutnes := 1000
	incrementar(quantidadeRoutnes)
	wg.Wait()
	fmt.Println("O valor do contador e ", contador)
}

func incrementar(goroutnes int) {
	wg.Add(goroutnes)

	for i := 0; i < goroutnes; i++ {
		go func() {
			defer mu.Unlock()
			defer wg.Done()

			mu.Lock()
			a := contador
			fmt.Println("O valor do contador e ", a)
			runtime.Gosched()
			a++
			contador = a
		}()
	}
}
